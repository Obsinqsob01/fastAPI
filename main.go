package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

/*-----------
 Data Types
-----------*/

type (
	task struct {
		ID int `json:"id"`
		Title string `json:"title"`
		Description string `json:"description"`
	}
)

/*-----------
 global vars
------------*/

var (
	tasks = map[int]*task{}
	seq = 1
)

/*--------
 Handlers
--------*/

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

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
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

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Routes
	e.POST("/tasks/", createTask)
	e.GET("/tasks/:id", getTask)
	e.GET("/tasks/", getTasks)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	// Run the server
	e.Logger.Fatal(e.Start(":8080"))
}
