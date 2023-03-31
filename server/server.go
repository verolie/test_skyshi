package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/verolie/test-skyshi/activity"
	"github.com/verolie/test-skyshi/migrate"
	"github.com/verolie/test-skyshi/todo"

	"github.com/labstack/echo/v4"
)

const (
	CONST_API_URL string = ":8090"
)

func RunServer() {
	migrate.Init()

	var apiUrl string
	e := echo.New()

	registerServer(e)

	checkURL := os.Getenv("API_URL")

	if checkURL != "" {
		apiUrl = checkURL
	} else {
		apiUrl = CONST_API_URL
	}

	fmt.Println("server started at port: " + apiUrl)
	e.Logger.Fatal(e.Start(apiUrl))
}

func registerServer(e *echo.Echo) {
	//activities
	e.GET("/activity-groups", ActivityHandler)
	e.POST("/activity-groups", ActivityHandler)
	e.GET("/activity-groups/:id", usersHandlerParam)
	e.PATCH("/activity-groups/:id", usersHandlerParam)
	e.DELETE("/activity-groups/:id", usersHandlerParam)

	//todo
	e.GET("/todo-items", TodoHandler)
	e.POST("/todo-items", TodoHandler)
	e.GET("/todo-items/:id", usersHandlerParamTodo)
	e.PATCH("/todo-items/:id", usersHandlerParamTodo)
	e.DELETE("/todo-items/:id", usersHandlerParamTodo)
}

func ActivityHandler(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodGet:
		return activity.GetActivities(c)
	case http.MethodPost:
		return activity.CreateActivity(c)
	default:
		return c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func usersHandlerParam(c echo.Context) error {
	switch c.Request().Method {

	case http.MethodGet:
		return activity.GetActivity(c)
	case http.MethodPatch:
		fmt.Println("patch")
		return activity.UpdateActivity(c)
	case http.MethodDelete:
		return activity.DeleteActivity(c)
	default:
		return c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func TodoHandler(c echo.Context) error {
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
		return todo.DeleteTodo(c)
	default:
		return c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}
