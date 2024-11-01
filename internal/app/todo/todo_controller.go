package todo

import (
	"go-boilerplate-api/internal/model"
	"go-boilerplate-api/internal/pkg/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
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
	defaultUser string
	service     TodoService
}

func ConstructController(db *db.Database) TodoController {
	return &todoController{
		defaultUser: viper.GetString("server.defaultUser"),
		service:     ConstructService(db),
	}
}

// Todo Api godoc
//
//	@Summary		CreateItem Todo Item
//	@Description	CreateItem Todo Item
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		201			{string}	CreateItem					todo	item
//	@Param			todoItem	body		CreateTodoRequest	true	"CreateItem Todo Item"
//	@Router			/todos [post]
func (t *todoController) CreateItem(ctx *gin.Context) {
	var creatingItem CreateTodoRequest
	if err := ctx.ShouldBindJSON(&creatingItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity := convertToEntity(creatingItem)
	createdItem, err := t.service.Create(t.defaultUser, entity)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, createdItem)
	}
}

func convertToEntity(dto CreateTodoRequest) *model.Todo {
	return &model.Todo{
		Description: dto.Description,
		StartDate:   (*time.Time)(dto.StartDate),
		EndDate:     (*time.Time)(dto.EndDate),
	}
}

// Todo Api godoc
//
//	@Summary		List All Todo Items
//	@Description	List All Todo Items
//	@Tags			todos
//	@Produce		json
//	@Param			userId	query		string	true	"the user id to filter"	example(henry.chou)
//	@Success		200	{string}	FindAll	todo	item
//	@Router			/todos [get]
func (t *todoController) FindAll(ctx *gin.Context) {
	userId := ctx.Query("userId")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}
	items, err := t.service.List(userId)
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
//	@Description	do ping
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	GetItem	todo	item
//	@Param			itemId	path		string	true	"get item by id"	example(7ae9c676-fc23-47a2-abc1-591ad2859b67)
//	@Router			/todos/{itemId} [get]
func (t *todoController) GetItem(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	item, err := t.service.Get(t.defaultUser, itemId)

	if err != nil {
		responseInternalServerError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, item)
	}
}

// Todo Api godoc
//
//	@Summary		UpdateItem Todo Item
//	@Description	UpdateItem Todo Item
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200			{string}	UpdateItem					todo	item
//	@Param			itemId		path		string						true	"the item id to be updated"	example(7d105cc8-a709-4a28-ae96-f0270bc5ad20)
//	@Param			todoItem	body		UpdateTodoRequest	true	"UpdateItem Todo Item"
//	@Router			/todos/{itemId} [patch]
func (t *todoController) UpdateItem(ctx *gin.Context) {
	var updatingItem UpdateTodoRequest
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

	updatedItem, err := t.service.Update(t.defaultUser, model)

	if err != nil {
		responseInternalServerError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, updatedItem)
	}
}

// Todo Api godoc
//
//	@Summary		DeleteItem Todo Item
//	@Description	DeleteItem Todo Item
//	@Tags			todos
//	@Produce		json
//	@Success		204		{string}	DeleteItem	todo	item
//	@Param			itemId	path		string		true	"the item id to be deleted"	example(7d105cc8-a709-4a28-ae96-f0270bc5ad20)
//	@Router			/todos/{itemId} [delete]
func (t *todoController) DeleteItem(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	err := t.service.Delete(t.defaultUser, itemId)

	if err != nil {
		responseInternalServerError(ctx, err)
	} else {
		ctx.Status(http.StatusNoContent)
	}
}
