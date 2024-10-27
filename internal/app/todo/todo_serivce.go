package todo

import (
	todoEntity "go-boilerplate-api/internal/app/todo/entities"
	"go-boilerplate-api/internal/pkg/db"
)

type TodoService interface {
	Create(todo *todoEntity.Todo) error
	Update(id string, todo *todoEntity.Todo) error
	List()
	Get(id string)
	Delete(id string)
}

type todoService struct {
	database *db.Database
}

func ConstructService(db *db.Database) TodoService {
	return &todoService{database: db}
}

func (t *todoService) Create(todo *todoEntity.Todo) error {
	return nil
}

func (t *todoService) List() {}

func (t *todoService) Get(id string) {}

func (t *todoService) Update(id string, todo *todoEntity.Todo) error {
	return nil
}

// Delete implements TodoService.
func (t *todoService) Delete(description string) {
	panic("unimplemented")
}
