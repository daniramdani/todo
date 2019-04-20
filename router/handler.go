package router

import (
	"github.com/daniramdani/todo/config"
	
	todoQuery "github.com/daniramdani/todo/modules/todo/query"
	todoUseCase "github.com/daniramdani/todo/modules/todo/usecase"
)

// Service handler data structure
type Service struct {
	TodoUseCase      todoUseCase.TodoUseCase
}

// MakeHandler - function for initializing handler of the services
func MakeHandler() *Service {
	// initiate database connection here
	readDB := config.ReadPostgresDB()
	writeDB := config.WritePostgresDB()
	
	todoQuery := todoQuery.NewTodoQueryPostgre(readDB, writeDB)
	dUseCase := todoUseCase.NewTodoUseCase(todoQuery)

	return &Service{
		TodoUseCase:      dUseCase,
	}
}
