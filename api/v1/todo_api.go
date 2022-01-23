package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo/model"
	"todo/service"
)

type TodoApi struct{}

func (t *TodoApi) CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := model.TodoModel{
		Title:     c.PostForm("title"),
		Completed: completed,
	}
	err := service.TodoServiceApp.AddTodoItem(todo)
	if err != nil {
		panic(fmt.Sprintf("Save new item failed: %v", err.Error()))
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Todo item created successfully!",
		"resourceId": todo.ID,
	})
}

func (t *TodoApi) FetchAllTodos(c *gin.Context) {
	todos, err := service.TodoServiceApp.FetchAllTodoItems()
	if err != nil {
		panic("Fetch all items failed")
	}
	if len(todos) < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No todo found!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   todos,
		})
	}
}

func (t *TodoApi) FetchSingleTodo(c *gin.Context) {
	id := c.Param("id")
	item, err := service.TodoServiceApp.FetchTodoItem(id)
	if err != nil {
		panic(fmt.Sprintf("Fetch item %v faild: %v", id, err.Error()))
	}
}
