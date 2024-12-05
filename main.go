package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/anupongpk/todo-go-gin/auth"
	"github.com/anupongpk/todo-go-gin/todo"
)

func main() {
	// connect databse
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// auto migrate
	db.AutoMigrate(&todo.Todo{})

	// Gin Routes
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// routes Auth
	r.GET("/tokenz", auth.AccessToken("===signature==="))
	protected := r.Group("", auth.Protect([]byte("===signature===")))

	handler := todo.NewTodoHandler(db)
	protected.POST("/todos", handler.NewTask)

	// Run
	r.Run()
}
