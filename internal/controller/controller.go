package controller

import (
	"github.com/gin-gonic/gin"
	model "github.com/jaxchow/gin/internal/model"
	// passport "github.com/jaxchow/planner-task/pkg/Passport"
)

func InitController(http *gin.RouterGroup) {
	model.ConnectDatabase()
	// apiRoute := http
	// model.RegisterModel(model.DB, &todo.Todo{})
	// todo.RestRoute(apiRoute.Group("/todos"))
	// passport.RestRoute(apiRoute.Group("/passport"))

}
