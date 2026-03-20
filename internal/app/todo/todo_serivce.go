package todo

import (
	"context"
	"errors"
	"go-boilerplate-api/internal/app/model"
	"go-boilerplate-api/internal/pkg/log"

	"github.com/google/uuid"
)

type TodoService struct {
	todoRepository TodoRepository
}

func NewTodoService(repository TodoRepository) TodoService {
	return TodoService{todoRepository: repository}
}

func (t *TodoService) Create(ctx context.Context, userId string, todo *model.Todo) (*model.Todo, error) {
	if todo == nil {
		message := "todo cannot be nil"
		log.Error(message)
		return nil, errors.New(message)
	}
	todo.UserId = userId
	return t.todoRepository.Create(ctx, todo)
}

func (t *TodoService) List(ctx context.Context, userId string) ([]*model.Todo, error) {
	items := []*model.Todo{}
	if userId == "" {
		return items, errors.New("userId is required to list todo items")
	}
	return t.todoRepository.List(ctx, userId)
}

func (t *TodoService) Get(ctx context.Context, userId string, rawItemId string) (*model.Todo, error) {
	itemId, err := uuid.Parse(rawItemId)
	if err != nil {
		return nil, err
	}
	return t.todoRepository.Get(ctx, userId, itemId)
}

func (t *TodoService) Update(ctx context.Context, userId string, todo *model.Todo) (*model.Todo, error) {
	return t.todoRepository.Update(ctx, userId, todo)
}

func (t *TodoService) Delete(ctx context.Context, userId string, rawItemId string) error {
	itemId, err := uuid.Parse(rawItemId)
	if err != nil {
		return err
	}
	return t.todoRepository.Delete(ctx, userId, itemId)
}
