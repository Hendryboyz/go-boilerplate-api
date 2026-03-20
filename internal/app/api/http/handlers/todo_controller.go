package handlers

import (
	"go-boilerplate-api/internal/app/api/http/dto"
	"go-boilerplate-api/internal/app/model"
	"go-boilerplate-api/internal/app/todo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoController struct {
	defaultUser string
	service     todo.TodoService
}

func NewTodoController(
	todoService todo.TodoService,
) *TodoController {
	return &TodoController{
		defaultUser: "Henry",
		service:     todoService,
	}
}

// Todo Api godoc
//
//	@Summary		Create Todo Item
//	@Description	Create a new todo item for the user
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			todoItem	body		dto.CreateTodoRequest	true	"Create Todo Item"
//	@Success		201			{object}	model.Todo
//	@Failure		400			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/todo [post]
func (t *TodoController) CreateItem(ctx *gin.Context) {
	var creatingItem dto.CreateTodoRequest
	if err := ctx.ShouldBindJSON(&creatingItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity := convertToEntity(creatingItem)
	createdItem, err := t.service.Create(ctx, t.defaultUser, entity)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, createdItem)
	}
}

func convertToEntity(dto dto.CreateTodoRequest) *model.Todo {
	return &model.Todo{
		Description: dto.Description,
		StartDate:   (*time.Time)(dto.StartDate),
		EndDate:     (*time.Time)(dto.EndDate),
	}
}

// Todo Api godoc
//
//	@Summary		List All Todo Items
//	@Description	List all todo items belonging to the specified user
//	@Tags			todos
//	@Produce		json
//	@Param			userId	query		string	true	"the user id to filter"	example(henry.chou)
//	@Success		200		{array}		model.Todo
//	@Failure		400		{object}	map[string]any
//	@Failure		500		{object}	map[string]any
//	@Router			/todo [get]
func (t *TodoController) FindAll(ctx *gin.Context) {
	userId := ctx.Query("userId")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}
	items, err := t.service.List(ctx, userId)
	if err != nil {
		responseInternalServerError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, items)
	}
}

func responseInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
}

// Todo Api godoc
//
//	@Summary		Get Todo Item By Id
//	@Description	Get a single todo item by its ID
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			itemId	path		string	true	"todo item id"	example(7ae9c676-fc23-47a2-abc1-591ad2859b67)
//	@Success		200		{object}	model.Todo
//	@Failure		500		{object}	map[string]any
//	@Router			/todo/{itemId} [get]
func (t *TodoController) GetItem(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	item, err := t.service.Get(ctx, t.defaultUser, itemId)

	if err != nil {
		responseInternalServerError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, item)
	}
}

// Todo Api godoc
//
//	@Summary		Update Todo Item
//	@Description	Update an existing todo item by its ID
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			itemId		path		string					true	"todo item id"	example(7d105cc8-a709-4a28-ae96-f0270bc5ad20)
//	@Param			todoItem	body		dto.UpdateTodoRequest	true	"Update Todo Item"
//	@Success		200			{object}	model.Todo
//	@Failure		400			{object}	map[string]any
//	@Failure		500			{object}	map[string]any
//	@Router			/todo/{itemId} [patch]
func (t *TodoController) UpdateItem(ctx *gin.Context) {
	var updatingItem dto.UpdateTodoRequest
	itemId := ctx.Param("itemId")
	if err := ctx.ShouldBindJSON(&updatingItem); err != nil {
		responseInternalServerError(ctx, err)
		return
	}

	model := &model.Todo{
		ID:          uuid.MustParse(itemId),
		UserId:      t.defaultUser,
		Description: updatingItem.Description,
	}

	if updatingItem.StartDate != nil {
		model.StartDate = (*time.Time)(updatingItem.StartDate)
	}

	if updatingItem.EndDate != nil {
		model.EndDate = (*time.Time)(updatingItem.EndDate)
	}

	updatedItem, err := t.service.Update(ctx, t.defaultUser, model)

	if err != nil {
		responseInternalServerError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, updatedItem)
	}
}

// Todo Api godoc
//
//	@Summary		Delete Todo Item
//	@Description	Delete a todo item by its ID
//	@Tags			todos
//	@Produce		json
//	@Param			itemId	path	string	true	"todo item id"	example(7d105cc8-a709-4a28-ae96-f0270bc5ad20)
//	@Success		204
//	@Failure		500	{object}	map[string]any
//	@Router			/todo/{itemId} [delete]
func (t *TodoController) DeleteItem(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	err := t.service.Delete(ctx, t.defaultUser, itemId)

	if err != nil {
		responseInternalServerError(ctx, err)
	} else {
		ctx.Status(http.StatusNoContent)
	}
}
