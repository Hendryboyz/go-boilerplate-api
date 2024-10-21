package todo

import "go-boilerplate-api/db"

type TodoService interface {
	GetItem(id string)
	Delete(description string)
}

type todoService struct {
	database *db.Database
}

func ConstructService(db *db.Database) TodoService {
	return &todoService{database: db}
}

func (t *todoService) GetItem(id string) {}

// Delete implements TodoService.
func (t *todoService) Delete(description string) {
	panic("unimplemented")
}
