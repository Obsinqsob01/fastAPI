package main

import (
	"strconv"
	"github.com/labstack/echo"
	"net/http"
)

type Task struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

var (
	seq uint = 1
	tasks map[uint]Task
)

func AddTask(ctx echo.Context) error {
	task := Task{}

	err := ctx.Bind(&task)
	if err != nil {
		panic(err)
	}

	task.ID = seq

	tasks[seq] = task

	seq++

	return ctx.JSON(http.StatusOK, task)
}

func UpdateTask(ctx echo.Context) error {
	u64, err := strconv.ParseUint(ctx.FormValue("id"), 64, 10)
	if err != nil {
		panic(err)
	}

	task := tasks[uint(u64)]

	task.Title = ctx.FormValue("title")
	task.Description = ctx.FormValue("description")

	tasks[uint(u64)] = task

	return ctx.JSON(http.StatusOK, task)
}

func DeleteTask(ctx echo.Context) error {
	u64, err := strconv.ParseUint(ctx.FormValue("id"), 64, 10)
	if err != nil {
		panic(err)
	}

	delete(tasks, uint(u64))

	return ctx.NoContent(http.StatusOK)
}

func GetTasks(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, tasks)
}

func GetTask(ctx echo.Context) error {
	u64, err := strconv.ParseUint(ctx.Param("id"), 64, 10)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(http.StatusOK, tasks[uint(u64)])
}

func main() {
	e := echo.New()

	e.POST("/task/add", AddTask)
	e.PUT("/task/update", UpdateTask)
	e.DELETE("/task/delete", DeleteTask)
	e.GET("/task/get/all", GetTasks)
	e.GET("/task/get/:id", GetTask)

	e.Logger.Fatal(e.Start(":8080"))
}
