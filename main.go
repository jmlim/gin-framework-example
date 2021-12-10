package main

import (
	"gin-framework-example/app/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// user router api start
	api.UserRouter(router)
	v1 := router.Group("/api")
	api.GetUsers(v1.Group("/page-test"))
	router.Run()
}
