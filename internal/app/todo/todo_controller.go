package todo

import (
	todoDto "go-boilerplate-api/internal/app/todo/dto"
	"go-boilerplate-api/internal/pkg/db"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	Create(*gin.Context)
	FindAll(*gin.Context)
	GetItem(*gin.Context)
	Delete(*gin.Context)
}

//	@BasePath	/v1
type todoController struct {
	service TodoService
}

func ConstructController(db *db.Database) TodoController {
	return &todoController{service: ConstructService(db)}
}

// PingExample godoc
//	@Summary		ping example
//	@Description	do ping
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		201			{string}	Create						todo	item
//	@Param			todoItem	body		todoDto.TodoCreatedRequest	true	"Create Todo Item"
//	@Router			/todos [post]
func (t *todoController) Create(ctx *gin.Context) {
	var creatingItem todoDto.TodoCreatedRequest
	if err := ctx.ShouldBindJSON(&creatingItem); err != nil {
		return
	}
}

// PingExample godoc
//	@Summary		ping example
//	@Description	do ping
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Create	todo	item
//	@Router			/todos [get]
func (t *todoController) FindAll(ctx *gin.Context) {
	panic("unimplemented")
}

// PingExample godoc
//	@Summary		ping example
//	@Description	do ping
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		201		{string}	Create	todo	item
//	@Param			itemId	path		string	true	"7ae9c676-fc23-47a2-abc1-591ad2859b67"
//	@Router			/todos/{itemId} [get]
func (t *todoController) GetItem(ctx *gin.Context) {
	panic("unimplemented")
}

// PingExample godoc
//	@Summary		ping example
//	@Description	do ping
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		201		{string}	Create	todo	item
//	@Param			itemId	path		string	true	"the item id to be deleted"	example(7d105cc8-a709-4a28-ae96-f0270bc5ad20)
//	@Router			/todos/{itemId} [delete]
func (t *todoController) Delete(ctx *gin.Context) {
	panic("unimplemented")
}
