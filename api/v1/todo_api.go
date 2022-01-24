package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"todo/common/response"
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
	response.Success(c, "Todo item created successfully!", gin.H{"resourceId": todo.ID})
}

func (t *TodoApi) FetchAllTodos(c *gin.Context) {
	todos, err := service.TodoServiceApp.FetchAllTodoItems()
	if err != nil {
		panic("Fetch all items failed")
	}
	if len(todos) < 1 {
		response.NotFound(c, "No todo found!", nil)
	} else {
		response.Success(c, "Todos fetched!", todos)
	}
}

func (t *TodoApi) FetchSingleTodo(c *gin.Context) {
	id := c.Param("id")
	item, err := service.TodoServiceApp.FetchTodoItem(id)
	if err != nil {
		panic(fmt.Sprintf("Fetch item %v faild: %v", id, err.Error()))
	}
	response.Success(c, "Todo found!", item)
}
