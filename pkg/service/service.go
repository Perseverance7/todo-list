package service

import (
	"github.com/Eagoker/todo-list"
	"github.com/Eagoker/todo-list/pkg/repository"
)

type Authorization interface{
	CreateUser(user todo.User) (int, error)
}

type TodoList interface{

}

type TodoItem interface{

}

type Service struct{
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}