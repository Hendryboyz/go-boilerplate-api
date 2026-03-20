package repository

import (
	"context"
	"errors"
	"fmt"
	"go-boilerplate-api/internal/app/model"
	"go-boilerplate-api/internal/pkg/log"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type GormTodoRepository struct {
	Db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *GormTodoRepository {
	return &GormTodoRepository{
		Db: db,
	}
}

func (t *GormTodoRepository) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	result := t.Db.WithContext(ctx).Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return todo, nil
}

func (t *GormTodoRepository) List(ctx context.Context, userId string) ([]*model.Todo, error) {
	items := []*model.Todo{}
	if userId == "" {
		return items, errors.New("userId is required to list todo items")
	}
	if err := t.Db.WithContext(ctx).Where("user_id = ?", userId).Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (t *GormTodoRepository) Get(ctx context.Context, userId string, itemId uuid.UUID) (*model.Todo, error) {
	var item *model.Todo
	err := t.Db.WithContext(ctx).First(&item, "user_id = ? AND id = ?", userId, userId).Error
	return item, err
}

func (t *GormTodoRepository) Update(ctx context.Context, userId string, todo *model.Todo) (*model.Todo, error) {
	existingTodo := &model.Todo{}
	if err := t.Db.WithContext(ctx).First(&existingTodo, "id = ? AND user_id = ?", todo.ID, userId).Error; err != nil {
		return nil, err
	}

	updatesFields := map[string]any{}
	if todo.Description != "" {
		updatesFields["description"] = todo.Description
	}
	if todo.StartDate != nil {
		updatesFields["start_date"] = todo.StartDate
	}
	if todo.EndDate != nil {
		updatesFields["end_date"] = todo.EndDate
	}
	tx := t.Db.WithContext(ctx).Model(&existingTodo).Updates(updatesFields)
	if err := tx.Error; err != nil {
		return nil, err
	} else {
		return todo, nil
	}
}

func (t *GormTodoRepository) Delete(ctx context.Context, userId string, itemId uuid.UUID) error {
	err := t.Db.WithContext(ctx).Where("user_id = ? AND id = ?", userId, itemId).Delete(&model.Todo{}).Error
	if err != nil {
		log.Error(fmt.Sprintf("fail to delete todo item[%s] for user[%s]", itemId, userId))
		return err
	} else {
		return nil
	}
}
