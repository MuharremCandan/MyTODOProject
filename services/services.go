package services

import (
	"database/sql"
	"net/http"
	"strconv"
	"todo-project/config"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	Id        int    `json:"Id"`
	Detail    string `json:"Detail"`
	Completed bool   `json:"COmpleted"`
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Api is working, hi :) ")
}

func AllTodos(c echo.Context) error {
	db, err := config.GetDB()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	defer db.Close()

	rows, _ := db.Query("Select Id ,Detail , Completed from Todos")

	defer rows.Close()

	todos := []Todo{}

	for rows.Next() {
		todoItem := Todo{}
		rows.Scan(&todoItem.Id, &todoItem.Detail, &todoItem.Completed)
		todos = append(todos, todoItem)
	}

	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {

	db, _ := config.GetDB()
	defer db.Close()

	u := Todo{}
	if err := c.Bind(u); err != nil {
		return err
	}

	statement, _ := db.Prepare("Insert into Todos(Detail ,Completed) Values (?,?)")
	statement.Exec(u.Detail, u.Completed)

	defer statement.Close()

	return c.JSON(http.StatusCreated, u)

}

func GetTodo(c echo.Context) error {

	db, _ := config.GetDB()
	defer db.Close()

	id, _ := strconv.Atoi(c.Param("id"))

	var todo Todo
	statement, _ := db.Prepare("SELECT Id, Detail, Completed FROM Todos WHERE Id = ?")
	err := statement.QueryRow(id).Scan(&todo.Id, &todo.Detail, &todo.Completed)

	defer statement.Close()

	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func UpdateTodoIsComplete(c echo.Context) error {
	db, _ := config.GetDB()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err.Error())
	}

	statement, _ := db.Prepare("UPDATE Todos SET Completed = 1 Where Id = ?")

	_, err = statement.Exec(id)
	defer statement.Close()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func UpdateTodoIsUncompleted(c echo.Context) error {

	db, _ := config.GetDB()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err.Error())
	}
	statement, _ := db.Prepare("UPDATE Todos SET Completed = 0 Where Id = ?")
	_, err = statement.Exec(id)
	defer statement.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)

}

func DeleteTodo(c echo.Context) error {
	db, _ := config.GetDB()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err.Error())
	}

	statement, _ := db.Prepare("delete from todos where Id = ?")
	_, err = statement.Exec(id)
	defer statement.Close()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)

}
