package todo

type TodoService interface {
	GetItem(id string)
	Delete(description string)
}

type todoService struct{}

func ConstructService() TodoService {
	return &todoService{}
}

func (t *todoService) GetItem(id string) {}

// Delete implements TodoService.
func (t *todoService) Delete(description string) {
	panic("unimplemented")
}
