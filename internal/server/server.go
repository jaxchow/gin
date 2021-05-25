package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jaxchow/gin/docs"
	todo "github.com/jaxchow/gin/internal/todo"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type HttpServer struct {
	// port      uint
	engine    *gin.Engine
	rootRoute *gin.RouterGroup
	// apiRoute  *gin.RouterGroup
}

func NewServer() HttpServer {
	http := new(HttpServer)
	http.New()
	return *http
}

func (httpSever *HttpServer) GetEngine() *gin.Engine {
	return httpSever.engine
}

func (httpServer *HttpServer) New() *HttpServer {
	httpServer.engine = CreateHttpEngine()
	apiRoute := httpServer.engine.Group("/api/v1")
	todosRoute := apiRoute.Group("/todos")
	todosRoute.GET("", todo.FindTodos)
	httpServer.SwaggerApi()
	// controller.InitController(httpServer.rootRoute)
	httpServer.StartHTTP()
	return httpServer
}

func (httpServer *HttpServer) StartHTTP() {
	httpServer.rootRoute = httpServer.engine.Group("/")
	// httpServer.apiRoute = httpServer.rootRoute.Group("/api/v1")
	// controller.InitController(httpServer.apiRoute)
	// todosRoute := httpServer.apiRoute.Group("/todos")
	// todosRoute.GET("", todo.FindTodos)
	httpServer.engine.Run()
}

// func (httpServer *HttpServer) RegisterRoute(){
// 	httpServer.apiRoute.
// }

func (httpServer *HttpServer) SwaggerApi() {
	httpServer.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func CreateHttpEngine() *gin.Engine {
	engine := gin.Default()

	return engine
}

// func (httpServer *HttpServer) CreateAPIGroup(path string) *gin.RouterGroup {
// 	route := httpServer.apiRoute.Group(path)
// 	return route
// }
