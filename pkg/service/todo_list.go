package service

import (
	todo "github.com/Glebegor/do-app"
	"github.com/Glebegor/do-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}
func (s *TodoListService) Create(UserId int, list todo.TodoList) (int, error) {
	return s.repo.Create(UserId, list)
}
func (s *TodoListService) GetAll(UserId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(UserId)
}
func (s *TodoListService) GetById(UserId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(UserId, listId)
}
func (s *TodoListService) Delete(UserId, listId int) error {
	return s.repo.Delete(UserId, listId)
}
func (s *TodoListService) Update(UserId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(UserId, listId, input)
}
