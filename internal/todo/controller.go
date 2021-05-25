package Todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获todos取
// @Description 查询指定数据
// @Tags todos
// @Accept json
// @Param limit query number false "10"
// @Param offset query number false "0"
// @Success 200 {string} json "{"data": []}"
// @Failure 400 {string} json "{"error": "Record not found!"}"
// @Router /todos [GET]

func FindTodos(c *gin.Context) {
	var todos []Todo
	limit, _ := strconv.ParseInt(c.Param("limit"), 10, 8)
	offset, _ := strconv.ParseInt(c.Param("offset"), 10, 8)
	todos = FindAll(QueryTodoInput{Offset: offset, Limit: limit})
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// @Summary 获取指定单一book
// @Description 获取指定ID book
// @Tags todos
// @Accept json
// @Param id path integer true "bookID"
// @Success 200 {string} json "{"data": "{}"}"
// @Failure 400 {string} json "{"error": "who are you"}"
// @Router /todos/:id [GET]

func FindTodo(c *gin.Context) {

	var todo Todo

	id := c.Param("id")
	todo, err := FindOne(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not fund!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// @Summary 创建book
// @Description  创建book对象
// @Tags todos
// @Accept json
// @Param id path integer true "bookID"
// @Param title body model.book true "Add group"
// @Success 201 {string} json "{"data": "{}"}"
// @Failure 400 {string} json "{"error": "err.Error()"}"
// @Router /todos/:id [POST]

func CreateTodo(c *gin.Context) {
	var input Todo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := Todo{}
	CreateOne(&todo)

	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

// @Summary 更新book对象
// @Description  更新指定book对象属性值
// @Tags todos
// @Accept json
// @Param id path integer true "bookID"
// @Success 200 {string} json "{"data": "{}"}"
// @Failure 400 {string} json "{"error": "err.Error()"}"
// @Router /todos/:id [PUT]

func UpdateTodo(c *gin.Context) {
	var todo Todo
	var input Todo
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, _ = UpdateOne(id, input)
	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

// @Summary 删除指定book
// @Description  删除指定ID的book对象
// @Tags todos
// @Accept json
// @Param id path integer true "bookID"
// @Success 204 {string} string "{"data": "{}"}"
// @Failure 400 {string} string "{"error": "Record not found!"}"
// @Router /todos/:id [DELETE]

func DeleteTodo(c *gin.Context) {
	// var todo Todo
	id := c.Param("id")
	_, _ = DeleteOne(id)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func RestRoute(route *gin.RouterGroup) {
	route.GET("", FindTodos)
	route.GET("/:id", FindTodo)
	route.POST("", CreateTodo)
	route.PATCH("/:id", UpdateTodo)
	route.DELETE("/:id", DeleteTodo)
}
