package router

import (
	"fmt"
	"os"
	"strconv"
	todoPresenter "github.com/daniramdani/todo/modules/todo/presenter"
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
)

const DefaultPort = 8080

// HTTPServerMain - function for initializing main HTTP server
func (s *Service) HTTPServerMain() {

	e := echo.New()
	e.Use(em.CORSWithConfig(em.CORSConfig{
		AllowMethods: []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	if os.Getenv("DEBUG") == "1" {
		e.Debug = true
	}

	
	todoHandler := todoPresenter.NewTodoHTTPHandler(s.TodoUseCase)
	todoGroup := e.Group("/v1/todo")
	todoHandler.Mount(todoGroup)
	
	// set REST port
	var port uint16
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			port = DefaultPort
		} else {
			port = uint16(portInt)
		}
	} else {
		port = DefaultPort
	}

	listenerPort := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(listenerPort))

}
