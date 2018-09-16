package main

import (
	"strconv"
	"github.com/labstack/echo"
	"net/http"
)

type Task struct {
	id uint `json:"id"`
	title string `json:"title"`
	description string `json:"description"`
}

var (
	seq uint = 1
	tasks map[uint]Task
)

func addTask(ctx echo.Context) error {
	task := Task{}

	err := ctx.Bind(&task)
	if err != nil {
		panic(err)
	}

	task.id = seq

	tasks[seq] = task

	seq++

	return ctx.JSON(http.StatusOK, task)
}

func updateTask(ctx echo.Context) error {
	u64, err := strconv.ParseUint(ctx.FormValue("id"), 64, 10)
	if err != nil {
		panic(err)
	}

	task := tasks[uint(u64)]

	task.title = ctx.FormValue("title")
	task.description = ctx.FormValue("description")

	tasks[uint(u64)] = task

	return ctx.JSON(http.StatusOK, task)
}

func deleteTask(ctx echo.Context) error {
	u64, err := strconv.ParseUint(ctx.FormValue("id"), 64, 10)
	if err != nil {
		panic(err)
	}

	delete(tasks, uint(u64))

	return ctx.NoContent(http.StatusOK)
}

func getTasks(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, tasks)
}

func getTask(ctx echo.Context) error {
	u64, err := strconv.ParseUint(ctx.Param("id"), 64, 10)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(http.StatusOK, tasks[uint(u64)])
}

func main() {
	e := echo.New()

	e.POST("/task/add", addTask)
	e.PUT("/task/update", updateTask)
	e.DELETE("/task/delete", deleteTask)
	e.GET("/task/get/all", getTasks)
	e.GET("/task/get/:id", getTask)

	e.Logger.Fatal(e.Start(":8080"))
}
