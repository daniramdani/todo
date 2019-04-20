package presenter

import (
	"github.com/daniramdani/todo/helper"
  "github.com/daniramdani/todo/modules/todo/usecase"
  "github.com/daniramdani/todo/modules/todo/sanitizer"
	"github.com/daniramdani/todo/state"
	"github.com/labstack/echo"
  "net/http"
  "strconv"
  //"fmt"
)

type todoHTTPHandler struct {
	TodoUseCase usecase.TodoUseCase
}

// NewTodoHTTPHandler - function for initializing profile http handler
func NewTodoHTTPHandler(todoUseCase usecase.TodoUseCase) *todoHTTPHandler {
	return &todoHTTPHandler{
		TodoUseCase: todoUseCase,
	}
}

// Mount - function for mounting route
func (h *todoHTTPHandler) Mount(group *echo.Group) {
  group.POST("", h.CreateTodo)
  group.GET("", h.GetAllTodo)
  group.GET("/:id", h.GetTodo)
  group.PUT("/:id", h.UpdateTodo)
  group.DELETE("/:id", h.DeleteTodo)
  group.PUT("/status/:id", h.UpdateStatusTodo)
}

// GetAllTodo - function for get all todo
func (h *todoHTTPHandler) GetAllTodo(c echo.Context) error {
	response := helper.TodoResponse()
	result, err := h.TodoUseCase.GetAllTodo()
	if err != nil && err != state.ErrDataNotFound {
		response.ErrorMessage = err.Error()
		response.Code = http.StatusInternalServerError
		return c.JSON(http.StatusInternalServerError, response)
	}
	response.Code = http.StatusOK
	response.Data = result
	return c.JSON(http.StatusOK, response)
}

// GetTodo - function for getting detail todo by id
func (h *todoHTTPHandler) GetTodo(c echo.Context) error {
  response := helper.TodoResponse()
  id, _ := strconv.Atoi(c.Param("id"))
	result, err := h.TodoUseCase.GetTodo(id)
	if err != nil && err != state.ErrDataNotFound {
		response.ErrorMessage = err.Error()
		response.Code = http.StatusInternalServerError
		return c.JSON(http.StatusInternalServerError, response)
	}
	response.Code = http.StatusOK
	response.Data = result
  return c.JSON(http.StatusOK, response)
}

// UpdateTodo - function for update todo by id
func (h *todoHTTPHandler) UpdateTodo(c echo.Context) error {
  id, _ := strconv.Atoi(c.Param("id"))
  todo, err := sanitizer.Todo(c)
	if err != nil && err != state.ErrDataNotFound {
    response := helper.TodoResponse()
		response.ErrorMessage = err.Error()
		response.Code = http.StatusInternalServerError
		return c.JSON(http.StatusInternalServerError, response)
  }
  response, err := h.TodoUseCase.UpdateTodo(todo, id)
	if err != nil && err != state.ErrDataNotFound {
		response.ErrorMessage = err.Error()
		response.Code = http.StatusInternalServerError
		return c.JSON(http.StatusInternalServerError, response)
  }
	response.Code = http.StatusOK
  return c.JSON(http.StatusOK, response)
}

// CreateTodo - function for create a todo
func (h *todoHTTPHandler) CreateTodo(c echo.Context) error {
  todo, err := sanitizer.Todo(c)
	if err != nil && err != state.ErrDataNotFound {
    response := helper.TodoResponse()
		response.ErrorMessage = err.Error()
		response.Code = http.StatusInternalServerError
		return c.JSON(http.StatusInternalServerError, response)
  }
  response, err := h.TodoUseCase.CreateTodo(todo)
	if err != nil && err != state.ErrDataNotFound {
		response.ErrorMessage = err.Error()
		response.Code = http.StatusInternalServerError
		return c.JSON(http.StatusInternalServerError, response)
  }
	response.Code = http.StatusOK
  return c.JSON(http.StatusOK, response)
}

// DeleteTodo - function for delete todo by id
func (h *todoHTTPHandler) DeleteTodo(c echo.Context) error {
  id, _ := strconv.Atoi(c.Param("id"))
  response, err := h.TodoUseCase.DeleteTodo(id)
	if err != nil && err != state.ErrDataNotFound {
		response.ErrorMessage = err.Error()
		response.Code = http.StatusInternalServerError
		return c.JSON(http.StatusInternalServerError, response)
  }
	response.Code = http.StatusResetContent
  return c.JSON(http.StatusOK, response)
}

// UpdateStatusTodo - function for update status todo by id
func (h *todoHTTPHandler) UpdateStatusTodo(c echo.Context) error {
  id, _ := strconv.Atoi(c.Param("id"))
  todo, err := sanitizer.Todo(c)
	if err != nil && err != state.ErrDataNotFound {
    response := helper.TodoResponse()
		response.ErrorMessage = err.Error()
		response.Code = http.StatusInternalServerError
		return c.JSON(http.StatusInternalServerError, response)
  }
  response, err := h.TodoUseCase.UpdateTodo(todo, id)
	if err != nil && err != state.ErrDataNotFound {
		response.ErrorMessage = err.Error()
		response.Code = http.StatusInternalServerError
		return c.JSON(http.StatusInternalServerError, response)
  }
	response.Code = http.StatusOK
  return c.JSON(http.StatusOK, response)
}