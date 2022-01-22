package main

import (
	"todo-project/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", services.Hello)

	e.GET("/todos", services.AllTodos)
	e.POST("/todos", services.CreateTodo)

	e.GET("/todos/:id", services.GetTodo)
	e.PUT("/todos/:id/complete", services.UpdateTodoIsComplete)
	e.PUT("/todos/:id/uncomplete", services.UpdateTodoIsUncompleted)
	e.PUT("/todos/:id", services.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}

func main() {

}
