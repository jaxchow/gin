package Todo

import (
	model "github.com/jaxchow/gin/internal/model"
)

type QueryTodoInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

// FindAll

func FindAll(query QueryTodoInput) []Todo {
	var todos []Todo
	model.DB.Find(&todos).Limit(query.Limit).Offset(query.Offset)
	return todos
}

// FindOne

func FindOne(id string) (Todo, error) {
	var todo Todo
	err := model.DB.First(&todo).Error
	if err != nil {
		return Todo{Title: "12112"}, nil
	}
	return Todo{Title: "12112"}, err
}

// CreateOne

func CreateOne(input interface{}) Todo {
	var todo Todo

	model.DB.Create(&input)

	return todo
}

// UpdateOne

func UpdateOne(id string, input ...interface{}) (Todo, error) {

	// var input UpdateTodoInput
	var todo Todo
	if err := model.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, err
	} else {
		model.DB.Model(&todo).Updates(input)
		return todo, nil
	}

	// models.DB.Model(&todo).Updates(input)
}

// Delete one

func DeleteOne(id string) (Todo, error) {
	var todo Todo
	if err := model.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, err
	} else {
		model.DB.Delete(&todo)
		return todo, nil
	}

}
