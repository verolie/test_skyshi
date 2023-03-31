package server

import (
	"fmt"
	"net/http"

	"github.com/verolie/test-skyshi/handler"
	"github.com/verolie/test-skyshi/todo"

	"github.com/labstack/echo/v4"
)

func RunServer() {

	e := echo.New()

	registerServer(e)

	e.Logger.Fatal(e.Start(":3030"))
}

func registerServer(e *echo.Echo) {
	//activities
	e.GET("/activity-groups", usersHandler)
	e.POST("/activity-groups", usersHandler)
	e.GET("/activity-groups/:id", usersHandlerParam)
	e.PATCH("/activity-groups/:id", usersHandlerParam)
	e.DELETE("/activity-groups/:id", usersHandlerParam)

	//todo
	e.GET("/todo-items", usersHandlerTodo)
	e.POST("/todo-items", usersHandlerTodo)
	e.GET("/todo-items/:id", usersHandlerParamTodo)
	e.PATCH("/todo-items/:id", usersHandlerParamTodo)
	e.DELETE("/todo-items/:id", usersHandlerParamTodo)
}

func usersHandler(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodGet:
		return handler.GetUsers(c)
	case http.MethodPost:
		return handler.CreateUser(c)
	default:
		return c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func usersHandlerParam(c echo.Context) error {
	switch c.Request().Method {

	case http.MethodGet:
		return handler.GetUser(c)
	case http.MethodPatch:
		fmt.Println("patch")
		return handler.UpdateUser(c)
	case http.MethodDelete:
		return handler.DeleteUser(c)
	default:
		return c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func usersHandlerTodo(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodGet:
		return todo.GetUsersTodo(c)
	case http.MethodPost:
		return todo.CreateUserTodo(c)
	default:
		return c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func usersHandlerParamTodo(c echo.Context) error {
	switch c.Request().Method {

	case http.MethodGet:
		return todo.GetUserTodo(c)
	case http.MethodPatch:
		fmt.Println("patch")
		return todo.UpdateUserTodo(c)
	case http.MethodDelete:
		return handler.DeleteUser(c)
	default:
		return c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}
