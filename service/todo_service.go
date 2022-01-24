package service

import (
	"todo/global"
	"todo/model"
)

type TodoService struct{}

var TodoServiceApp = new(TodoService)

func (s *TodoService) AddTodoItem(todo model.TodoModel) error {
	res := global.DB.Save(&todo)
	return res.Error
}

func (s *TodoService) FetchAllTodoItems() ([]model.TodoResponseModel, error) {
	var items []model.TodoModel
	var _items []model.TodoResponseModel
	result := global.DB.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(items) == 0 {
		return []model.TodoResponseModel{}, nil
	}
	// 对todo的属性做一些转换以构建更好的响应体
	for _, item := range items {
		var completed bool
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_items = append(_items, model.TodoResponseModel{ID: item.ID, Title: item.Title, Completed: completed})
	}
	return _items, nil
}

func (s *TodoService) FetchTodoItem(id string) (model.TodoResponseModel, error) {
	var item model.TodoModel
	var resp model.TodoResponseModel
	res := global.DB.First(&item, id)
	if res.Error != nil {
		return resp, res.Error
	}
	completed := false
	if item.Completed == 1 {
		completed = true
	}
	resp = model.TodoResponseModel{ID: item.ID, Title: item.Title, Completed: completed}
	return resp, nil
}
