package service

import (
	"github.com/gin-gonic/gin"
	"gobyexample/awesome-go/bazel/go-tutorial/handler/publish"
)

func SetUpRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/publish", publish.Publish)
	return engine
}
