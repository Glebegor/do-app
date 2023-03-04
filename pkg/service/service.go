package service

import (
	todo "github.com/Glebegor/do-app"
	"github.com/Glebegor/do-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(UserId int, list todo.TodoList) (int, error)
	GetAll(UserId int) ([]todo.TodoList, error)
	GetById(UserId, listId int) (todo.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
