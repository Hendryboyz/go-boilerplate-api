package todo

import (
	"context"
	"go-boilerplate-api/internal/app/model"

	"github.com/google/uuid"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	List(ctx context.Context, userId string) ([]*model.Todo, error)
	Get(ctx context.Context, userId string, itemId uuid.UUID) (*model.Todo, error)
	Update(ctx context.Context, userId string, todo *model.Todo) (*model.Todo, error)
	Delete(ctx context.Context, userId string, itemId uuid.UUID) error
}
