package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type (
	task struct {
		ID int `json:"id"`
		Title string `json:"title"`
		Description string `json:"description"`
	}
)

var (
	tasks = map[int]*task{}
	seq = 1
)

func createTask(c echo.Context) error {
	t := &task{
		ID: seq,
	}

	if err := c.Bind(t); err != nil {
		return err
	}

	tasks[t.ID] = t
	seq++

	return c.JSON(http.StatusCreated, t)
}

func getTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, tasks[id])
}

func updateTask(c echo.Context) error {
	t := new(task)

	if err := c.Bind(t); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	tasks[id].Title = t.Title
	tasks[id].Description = t.Description
	
	return c.JSON(http.StatusOK, tasks[id])
}

func deleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(tasks, id)

	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users/", createTask)
	e.GET("/users/:id", getTask)
	e.PUT("/users/:id", updateTask)
	e.DELETE("/users/:id", deleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
