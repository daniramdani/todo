package query

import "github.com/daniramdani/todo/modules/todo/model"

// TodoQuery - todo query interface abstraction
type TodoQuery interface {
	GetAllTodo() (*[]*model.Todo, error)
	GetTodo(id int) (*model.Todo, error)
	UpdateTodo(todo *model.Todo, id int) (*model.Todo, error)
	CreateTodo(todo *model.Todo) (*model.Todo, error)
	DeleteTodo(id int) (*model.Todo, error)
}
