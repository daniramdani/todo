package usecase

import "github.com/daniramdani/todo/modules/todo/model"
import "github.com/daniramdani/todo/helper"

// TodoUseCase - todo use case interface abstraction
type TodoUseCase interface {
	GetAllTodo() (*[]*model.Todo, error)
	GetTodo(id int) (*model.Todo, error)
	UpdateTodo(todo *model.Todo, id int) (*helper.TodoResponseStrc, error)
	CreateTodo(todo *model.Todo) (*helper.TodoResponseStrc, error)
	DeleteTodo(id int) (*helper.TodoResponseStrc, error)
}
