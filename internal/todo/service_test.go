package Todo_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/jaxchow/gin/internal/Todo"
	model "github.com/jaxchow/gin/internal/model"
)

func Init() {
	model.ConnectDatabase()
	model.RegisterModel(model.DB, &Todo.Todo{})
	model.DB.Begin()
}

func Reset() {
	model.DB.Rollback()
}

func TestFindAll(t *testing.T) {
	Init()
	var query Todo.QueryTodoInput
	query.Limit = 10
	query.Offset = 0
	todos := Todo.FindAll(query)
	assert.Equal(t, len(todos), 0)
	Reset()
}

func TestCreateOne(t *testing.T) {
	Init()
	var todo Todo.Todo
	var input Todo.Todo
	input.Author = "jaxchow"
	input.Title = "today meeting"
	input.Content = "大家一起聊聊"
	input.Owner = "jaxchow"
	input.Priority = "1"
	todo = Todo.CreateOne(&input)
	// println(todo.Title)
	assert.Equal(t, todo.Title, "today meeting")
	Reset()
}

func TestFindOne(t *testing.T) {
	Init()

	todo, err := Todo.FindOne("2")
	print(todo.Author)
	assert.Equal(t, todo, err)
	Reset()
}
