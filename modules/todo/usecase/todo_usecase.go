package usecase

import (
	"github.com/daniramdani/todo/modules/todo/model"
  "github.com/daniramdani/todo/modules/todo/query"
  "github.com/daniramdani/todo/helper"
  log "github.com/sirupsen/logrus"
)

// todoUseCaseImpl - todo use case implementation
type todoUseCaseImpl struct {
  todoQueryRead query.TodoQuery
  todoQueryWrite query.TodoQuery
}

// NewTodoUseCase - function for initializing user use case implementation
func NewTodoUseCase(todoQueryRead query.TodoQuery) TodoUseCase {
	return &todoUseCaseImpl{
		todoQueryRead: todoQueryRead,
	}
}

// GetAllTodo - function for get all todo
func (du *todoUseCaseImpl) GetAllTodo() (*[]*model.Todo, error) {
	return du.todoQueryRead.GetAllTodo()
}

// GetTodo - function for get todo
func (du *todoUseCaseImpl) GetTodo(id int) (*model.Todo, error) {
	return du.todoQueryRead.GetTodo(id)
}

// UpdateTodo - function for update todo
func (du *todoUseCaseImpl) UpdateTodo(todo *model.Todo, id int) (*helper.TodoResponseStrc, error) {
	ctx := "TodoUseCaseImpl-UpdateTodo"
  
  res, err := du.todoQueryRead.UpdateTodo(todo, id)
	if err != nil {
		helper.Capture(log.ErrorLevel, err, ctx, "UpdateTodo")
		return nil, err
	}

	response := helper.TodoResponse()
	response.RequestID = helper.RandomStringBase64(16)
	response.Data = model.Todo{
		ID:            res.ID,
		Title:         res.Title,
		Description:   res.Description,
		Status:        res.Status,
	}

	return &response, nil
}

// CreateTodo - function for create a todo
func (du *todoUseCaseImpl) CreateTodo(todo *model.Todo) (*helper.TodoResponseStrc, error) {
	ctx := "TodoUseCaseImpl-UpdateTodo"
         
  res, err := du.todoQueryRead.CreateTodo(todo)
	if err != nil {
		helper.Capture(log.ErrorLevel, err, ctx, "UpdateTodo")
		return nil, err
	}

	response := helper.TodoResponse()
	response.RequestID = helper.RandomStringBase64(16)
	response.Data = model.Todo{
		ID:            res.ID,
		Title:         res.Title,
		Description:   res.Description,
		Status:        res.Status,
	}

	return &response, nil
}

// DeleteTodo - function for update todo
func (du *todoUseCaseImpl) DeleteTodo(id int) (*helper.TodoResponseStrc, error) {
	ctx := "TodoUseCaseImpl-DeleteTodo"
  
  _, err := du.todoQueryRead.DeleteTodo(id)
	if err != nil {
		helper.Capture(log.ErrorLevel, err, ctx, "DeleteTodo")
		return nil, err
	}

	response := helper.TodoResponse()
	response.RequestID = helper.RandomStringBase64(16)
	response.Data = model.Todo{
		ID:            id,
	}

	return &response, nil
}
