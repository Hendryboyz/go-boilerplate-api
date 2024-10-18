package todo

import "github.com/gin-gonic/gin"

type TodoController interface {
	Create(*gin.Context)
	FindAll(*gin.Context)
	GetItem(*gin.Context)
	Delete(*gin.Context)
}

type todoController struct {
	service TodoService
}

func ConstructController() TodoController {
	return &todoController{service: ConstructService()}
}

func (t *todoController) Create(*gin.Context) {
	panic("unimplemented")
}

func (t *todoController) FindAll(*gin.Context) {
	panic("unimplemented")
}

func (t *todoController) GetItem(*gin.Context) {
	panic("unimplemented")
}

func (t *todoController) Delete(*gin.Context) {
	panic("unimplemented")
}
