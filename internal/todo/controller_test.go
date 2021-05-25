package Todo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	models "github.com/jaxchow/gin/internal/model"
	"github.com/jaxchow/gin/internal/server"
)

var httpServer server.HttpServer

func mockHttp() {
	httpServer = server.NewServer()
	models.DB.Begin()
}

func reset() {
	models.DB.Rollback()
}

func TestGetTodoController(t *testing.T) {
	mockHttp()
	engine := httpServer.GetEngine()
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todos", nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)
	reset()
}
