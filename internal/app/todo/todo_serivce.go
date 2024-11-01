package todo

import (
	"errors"
	"fmt"
	"go-boilerplate-api/internal/model"
	"go-boilerplate-api/internal/pkg/db"
	"go-boilerplate-api/internal/pkg/log"

	"github.com/google/uuid"
)

type TodoService interface {
	Create(userId string, todo *model.Todo) (*model.Todo, error)
	List(userId string) ([]*model.Todo, error)
	Get(userId string, itemId string) (*model.Todo, error)
	Update(userId string, todo *model.Todo) (*model.Todo, error)
	Delete(userId string, itemId string) error
}

type todoService struct {
	database *db.Database
}

func ConstructService(db *db.Database) TodoService {
	return &todoService{database: db}
}

func (t *todoService) Create(userId string, todo *model.Todo) (*model.Todo, error) {
	if todo == nil {
		message := "todo cannot be nil"
		log.Error(message)
		return nil, errors.New(message)
	}
	todo.UserId = userId
	result := t.database.Client.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return todo, nil
}

func (t *todoService) List(userId string) ([]*model.Todo, error) {
	items := []*model.Todo{}
	if userId == "" {
		return items, errors.New("userId is required to list todo items")
	}
	if err := t.database.Client.Where("user_id = ?", userId).Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (t *todoService) Get(userId string, itemId string) (*model.Todo, error) {
	var item *model.Todo
	uuid, err := uuid.Parse(itemId)
	if err != nil {
		return item, err
	}

	err = t.database.Client.First(&item, "user_id = ? AND id = ?", userId, uuid).Error
	return item, err
}

func (t *todoService) Update(userId string, todo *model.Todo) (*model.Todo, error) {
	existingTodo := &model.Todo{}
	if err := t.database.Client.First(&existingTodo, "id = ? AND user_id = ?", todo.ID, userId).Error; err != nil {
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
	tx := t.database.Client.Model(&existingTodo).Updates(updatesFields)
	if err := tx.Error; err != nil {
		return nil, err
	} else {
		return todo, nil
	}
}

func (t *todoService) Delete(userId string, itemId string) error {
	err := t.database.Client.Where("user_id = ? AND id = ?", userId, itemId).Delete(&model.Todo{}).Error
	if err != nil {
		log.Error(fmt.Sprintf("fail to delete todo item[%s] for user[%s]", itemId, userId))
		return err
	} else {
		return nil
	}
}
