package main

import (
	"gin-framework-example/app/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// user router api start
	api.UserRouter(router)
	router.Run()
}
