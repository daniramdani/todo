package sanitizer

import (
	"github.com/daniramdani/todo/modules/todo/model"
	"github.com/labstack/echo"
	"github.com/daniramdani/todo/helper"
	log "github.com/sirupsen/logrus"
)

func Todo(c echo.Context) (p *model.Todo, err error) {
	ctx := "TodoQuery-UpdateTodo"
	if err := c.Bind(&p); err != nil {
		helper.Capture(log.ErrorLevel, err, ctx, "bind_request")
		return nil, err
	}
	return p, nil
}
