package todo

import (
	"errors"
	todoEntity "go-boilerplate-api/internal/app/todo/entities"
	"go-boilerplate-api/internal/pkg/db"
	"go-boilerplate-api/internal/pkg/log"

	"github.com/google/uuid"
)

type TodoService interface {
	Create(todo *todoEntity.Todo) (*todoEntity.Todo, error)
	Update(id string, todo *todoEntity.Todo) error
	List(userId string) ([]*todoEntity.Todo, error)
	Get(id string) (*todoEntity.Todo, error)
	Delete(id string) error
}

type todoService struct {
	database *db.Database
}

func ConstructService(db *db.Database) TodoService {
	return &todoService{database: db}
}

func (t *todoService) Create(todo *todoEntity.Todo) (*todoEntity.Todo, error) {
	if todo == nil {
		message := "todo cannot be nil"
		log.Error(message)
		return nil, errors.New(message)
	}

	result := t.database.Client.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return todo, nil
}

func (t *todoService) List(userId string) ([]*todoEntity.Todo, error) {
	items := []*todoEntity.Todo{}
	if userId == "" {
		return items, errors.New("userId is required to list todo items")
	}
	if err := t.database.Client.Where("user_id = ?", userId).Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (t *todoService) Get(id string) (*todoEntity.Todo, error) {
	var item *todoEntity.Todo
	uuid, err := uuid.Parse(id)
	if err != nil {
		return item, err
	}

	err = t.database.Client.First(&item, "id = ?", uuid).Error
	return item, err
}

func (t *todoService) Update(id string, todo *todoEntity.Todo) error {
	return nil
}

func (t *todoService) Delete(description string) error {
	return nil
}
