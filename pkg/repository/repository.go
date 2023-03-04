package repository

import (
	todo "github.com/Glebegor/do-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(UserId int, list todo.TodoList) (int, error)
	GetAll(UserId int) ([]todo.TodoList, error)
	GetById(UserId, listId int) (todo.TodoList, error)
	Delete(UserId, listId int) error
	Update(UserId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(UserId, listId int) ([]todo.TodoItem, error)
	GetById(UserId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(UserId, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
