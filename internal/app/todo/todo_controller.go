package todo

import (
	todoDto "go-boilerplate-api/internal/app/todo/dto"
	"go-boilerplate-api/internal/pkg/db"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	CreateItem(*gin.Context)
	FindAll(*gin.Context)
	GetItem(*gin.Context)
	UpdateItem(*gin.Context)
	DeleteItem(*gin.Context)
}

// @BasePath	/v1
type todoController struct {
	service TodoService
}

func ConstructController(db *db.Database) TodoController {
	return &todoController{service: ConstructService(db)}
}

// Todo Api godoc
//
//	@Summary		CreateItem Todo Item
//	@Description	CreateItem Todo Item
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		201			{string}	CreateItem						todo	item
//	@Param			todoItem	body		todoDto.CreateTodoRequest	true	"CreateItem Todo Item"
//	@Router			/todos [post]
func (t *todoController) CreateItem(ctx *gin.Context) {
	var creatingItem todoDto.CreateTodoRequest
	if err := ctx.ShouldBindJSON(&creatingItem); err != nil {
		return
	}
}

// Todo Api godoc
//
//	@Summary		List All Todo Items
//	@Description	List All Todo Items
//	@Tags			todos
//	@Produce		json
//	@Success		200	{string}	FindAll	todo	item
//	@Router			/todos [get]
func (t *todoController) FindAll(ctx *gin.Context) {
	t.service.List()
}

// Todo Api godoc
//
//	@Summary		Get Todo Item By Id
//	@Description	do ping
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	GetItem	todo	item
//	@Param			itemId	path		string	true	"get item by id"	example(7ae9c676-fc23-47a2-abc1-591ad2859b67)
//	@Router			/todos/{itemId} [get]
func (t *todoController) GetItem(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	t.service.Get(itemId)
}

// Todo Api godoc
//
//	@Summary		UpdateItem Todo Item
//	@Description	UpdateItem Todo Item
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	UpdateItem	todo	item
//	@Param			itemId	path		string	true	"the item id to be updated"	example(7d105cc8-a709-4a28-ae96-f0270bc5ad20)
//	@Param			todoItem	body		todoDto.UpdateTodoRequest	true	"UpdateItem Todo Item"
//	@Router			/todos/{itemId} [patch]
func (t *todoController) UpdateItem(ctx *gin.Context) {
	var updatingItem todoDto.UpdateTodoRequest
	if err := ctx.ShouldBindJSON(&updatingItem); err != nil {
		return
	}
}

// Todo Api godoc
//
//	@Summary		DeleteItem Todo Item
//	@Description	DeleteItem Todo Item
//	@Tags			todos
//	@Produce		json
//	@Success		204		{string}	DeleteItem	todo	item
//	@Param			itemId	path		string	true	"the item id to be deleted"	example(7d105cc8-a709-4a28-ae96-f0270bc5ad20)
//	@Router			/todos/{itemId} [delete]
func (t *todoController) DeleteItem(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	t.service.Delete(itemId)
}
