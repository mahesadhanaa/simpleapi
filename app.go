package main

import (
	"database/sql"

	"github.com/mahesadhanaa/simpleapi/handlers"

	"github.com/mahesadhanaa/echo"
	_ "github.com/mahesadhanaa/go-sqlite3"
)

func main() {

	e := echo.New()
	db := initDB("storage.db")
	migrate(db)

	// daftar api
	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/tasks", handlers.PutTask(db))
	e.PUT("/tasks", handlers.EditTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		status INTEGER
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
