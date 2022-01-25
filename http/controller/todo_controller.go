package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo/common/response"
	"todo/model"
	"todo/service/todo"
)

type TodoApi struct{}

func (t *TodoApi) Store(c *gin.Context) {
	item := model.TodoModel{}
	err := c.ShouldBindJSON(&item)
	if err != nil {
		response.Fail(c, "Save new item failed: "+err.Error(), nil)
		return
	}
	var id uint
	id, err = todo.TodoServiceApp.Store(&item)
	if err != nil {
		response.Fail(c, "Save new item failed: "+err.Error(), nil)
		return
	}
	response.Success(c, "Todo item created successfully!", gin.H{"resourceId": id})
}

func (t *TodoApi) All(c *gin.Context) {
	todos, err := todo.TodoServiceApp.All()
	if err != nil {
		response.Fail(c, "Fetch all items failed :"+err.Error(), nil)
		return
	}
	if len(todos) < 1 {
		response.NotFound(c, "No todo found!", nil)
		return
	} else {
		response.Success(c, "Todos fetched!", todos)
	}
}

func (t *TodoApi) Show(c *gin.Context) {
	id := c.Param("id")
	item, err := todo.TodoServiceApp.Show(id)
	if err != nil {
		response.Fail(c, fmt.Sprintf("Fetch item %v faild: %v", id, err.Error()), nil)
		return
	}
	// Not found
	if item.ID == 0 {
		response.NotFound(c, "Item with id "+id+" not found!", nil)
		return
	}
	response.Success(c, "Todo found!", item)
}

func (t *TodoApi) Update(c *gin.Context) {
	//id := c.Param("id")

	item := model.TodoModel{}
	err := c.ShouldBindJSON(&item)
	if err != nil {
		response.Fail(c, "Update todo failed: "+err.Error(), nil)
		return
	}

}
