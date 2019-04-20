package query

import (
	"database/sql"
	"github.com/daniramdani/todo/helper"
	"github.com/daniramdani/todo/modules/todo/model"
	"github.com/daniramdani/todo/state"
	"github.com/jmoiron/sqlx"
  log "github.com/sirupsen/logrus"
  "fmt"
)

// todoQueryPostgre - todo query implementation
type todoQueryPostgre struct {
	dbRead  *sqlx.DB
	dbWrite *sqlx.DB
}

// NewTodoQueryPostgre - function for initializing todo query
func NewTodoQueryPostgre(dbRead *sqlx.DB, dbWrite *sqlx.DB) *todoQueryPostgre {
	return &todoQueryPostgre{dbRead: dbRead, dbWrite: dbWrite}
}

// GetAllTodo - function for get all todo
func (aq *todoQueryPostgre) GetAllTodo() (*[]*model.Todo, error) {
	ctx := "TodoQuery-GetAllTodo"

	sql := `SELECT id, title, description, status FROM todos`
	rs, err := aq.dbRead.Queryx(sql)
	if err != nil {
		helper.Capture(log.ErrorLevel, err, ctx, "query_database")
		return nil, err
	}
	defer rs.Close()

	var id int
	var title string
	var description string
	var status int

  var todos []*model.Todo

	for rs.Next() {
		err := rs.Scan(&id, &title, &description, &status)
		if err != nil {
			helper.Capture(log.ErrorLevel, err, ctx, "query_database")
			return nil, err
		}

		todo := &model.Todo{
			ID:               id,
			Title:        title,
			Description: description,
			Status:             status,
		}

		todos = append(todos, todo)
	}

	return &todos, nil
}

// GetAllTodo - function for get a todo
func (aq *todoQueryPostgre) GetTodo(id int) (*model.Todo, error) {
  todo := &model.Todo{}
  err := aq.dbRead.QueryRow(`SELECT id, title, description, status FROM todos WHERE id = $1`, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Status,
  )
  if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, state.ErrDataNotFound
		default:
			return nil, err
		}
	}
	
	return todo, nil
}

// UpdateTodo - function for update a todo
func (aq *todoQueryPostgre) UpdateTodo(todo *model.Todo, id int) (*model.Todo, error) {
  ctx := "TodoQueryPostgre-UpdateTodo"
  
	query := fmt.Sprintf(`UPDATE todos
                        SET title = $1, description = $2, status = $3
                        WHERE id = $4`)    
  
  if todo.Title == "" || todo.Description == "" {
    current, _ := aq.GetTodo(id)
    if todo.Title == "" {
      todo.Title = current.Title
    }
    if todo.Description == "" {
      todo.Description = current.Description
    }
  }

  _, err := aq.dbWrite.Exec(query,
    todo.Title,
    todo.Description,
    todo.Status,
    id,
  )
  if err != nil {
		helper.Capture(log.ErrorLevel, err, ctx, "execute_update_todo")
		return nil, err
  }
  
  res := &model.Todo{
		ID:            id,
		Title:         todo.Title,
		Description:   todo.Description,
		Status:        todo.Status,
	}
  
	return res, nil
}

// CreateTodo - function for create a todo
func (aq *todoQueryPostgre) CreateTodo(todo *model.Todo) (*model.Todo, error) {
  ctx := "TodoQueryPostgre-CreateTodo"
  
  id := 0
  err := aq.dbRead.QueryRow(`INSERT INTO todos (title, description, status) VALUES ($1, $2, $3) RETURNING id`, todo.Title, todo.Description, todo.Status).Scan(&id)
  if err != nil {
		helper.Capture(log.ErrorLevel, err, ctx, "execute_create_todo")
		return nil, err
  }
  
  res := &model.Todo{
    ID:            id,
		Title:         todo.Title,
		Description:   todo.Description,
		Status:        todo.Status,
	}
  
	return res, nil
}

// DeleteTodo - function for delete a todo
func (aq *todoQueryPostgre) DeleteTodo(id int) (*model.Todo, error) {
  ctx := "TodoQueryPostgre-DeleteTodo"
  
	query := fmt.Sprintf(`DELETE FROM todos WHERE id = $1`)              
  _, err := aq.dbWrite.Exec(query, id)
  if err != nil {
		helper.Capture(log.ErrorLevel, err, ctx, "execute_delete_todo")
		return nil, err
  }
  
  res := &model.Todo{
		ID:            id,
	}
  
	return res, nil
}