package integration

import (
	"bytes"
	"fmt"
	"github.com/daniramdani/todo/config"
  "github.com/daniramdani/todo/router"
  todoPresenter "github.com/daniramdani/todo/modules/todo/presenter"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"log"
	"net/http/httptest"
	"os"
	"strings"
)

var writeDB, readDB *sqlx.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("can't load .env file")
		os.Exit(2)
	}
	initDB()
}

func initDB() {
	if !strings.HasSuffix(os.Getenv("WRITE_DB_NAME"), "_test") || !strings.HasSuffix(os.Getenv("READ_DB_NAME"), "_test") {
		fmt.Println("Testing database must be suffixed with _test")
		os.Exit(1)
  }
  
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s",
  os.Getenv("WRITE_DB_USER"), os.Getenv("WRITE_DB_NAME"), os.Getenv("WRITE_DB_PASSWORD"), os.Getenv("WRITE_DB_HOST")))
	if err != nil {
		log.Fatalf("error connecting to DB: %s", err)
	}
	defer db.Close()

	db, err = sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s",
  os.Getenv("READ_DB_USER"), os.Getenv("READ_DB_NAME"), os.Getenv("READ_DB_PASSWORD"), os.Getenv("READ_DB_HOST")))
	if err != nil {
		log.Fatalf("error connecting to DB: %s", err)
	}
	defer db.Close()

	writeDB = config.WritePostgresDB()
  readDB = config.ReadPostgresDB()
  tx, _ := writeDB.Begin()
  tx.Exec(`DROP TABLE todos;`)
  err = tx.Commit()
	if err != nil {
		fmt.Println(err)
  }

  tx, _ = writeDB.Begin()
  tx.Exec(`CREATE TABLE todos(
            id serial PRIMARY KEY,
            title VARCHAR (50) NOT NULL,
            description TEXT NULL,
            status SMALLINT DEFAULT 0,
            created_at TIMESTAMP DEFAULT NOW(),
            modified_at TIMESTAMP NULL);`)
  err = tx.Commit()
	if err != nil {
		fmt.Println(err)
  }
}

// DoRequest simulate request http
func DoRequest(method, uri string, body *bytes.Buffer) (*httptest.ResponseRecorder, error) {
	resp := httptest.NewRecorder() //application/json
	req := httptest.NewRequest(method, uri, strings.NewReader(body.String()))

	e := echo.New()
	s := router.MakeHandler()

	// set the group of routes
	todoHandler := todoPresenter.NewTodoHTTPHandler(s.TodoUseCase)
	todoGroup := e.Group("/v1/todo")
	todoHandler.Mount(todoGroup)

	e.ServeHTTP(resp, req)

	return resp, nil
}