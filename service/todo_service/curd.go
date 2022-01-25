package todo_service

import (
	"errors"
	"todo/global"
	"todo/model"
)

type TodoService struct{}

var TodoServiceApp = new(TodoService)

func (s *TodoService) Store(title string, description string) error {
	item := model.TodoModel{
		Title:       title,
		Description: description,
	}
	res := global.DB.Save(&item)
	return res.Error
}

func (s *TodoService) Update(id int, title string, description string) error {
	res := global.DB.Model(&model.TodoModel{}).
		Where("id = ?", id).
		Update(model.TodoModel{Title: title, Description: description})
	if res.RowsAffected < 1 {
		return errors.New("record not found")
	}
	return res.Error
}

func (s *TodoService) All() ([]model.TodoResponseModel, error) {
	var items []model.TodoModel
	var resps []model.TodoResponseModel
	res := global.DB.Find(&items)
	if res.Error != nil {
		return nil, res.Error
	}
	if len(items) == 0 {
		return []model.TodoResponseModel{}, nil
	}
	// 对todo的属性做一些转换以构建更好的响应体
	for _, item := range items {
		resps = append(resps, model.TodoResponseModel{ID: item.ID, Title: item.Title, Completed: item.Completed})
	}
	return resps, nil
}

func (s *TodoService) Show(id string) (model.TodoResponseModel, error) {
	var item model.TodoModel
	var resp model.TodoResponseModel
	if err := global.DB.First(&item, id).Error; err != nil {
		return resp, err
	}
	resp = model.TodoResponseModel{ID: item.ID, Title: item.Title, Completed: item.Completed}
	return resp, nil
}
