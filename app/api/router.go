package api

import (
	"fmt"
	"gin-framework-example/app/service"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/users", func(c *gin.Context) {
		users := service.GetUsers(c)

		// 조회
		for key, value := range users {
			fmt.Println("key : ", key,
				", value.LastName: ", *value.LastName,
				", value.Address[0].City : ", *value.Address[0].City,
				", value.PaymentMethod[0].PaymentToken", *value.PaymentMethod[0].PaymentToken)
		}

		c.JSON(200, users)
	})

	/**
	ex: 	http://localhost:5000/users-email/kylebanker2@gmail.com
	*/
	router.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		users := service.GetUsersByName(c, name)

		// 조회
		for key, value := range users {
			fmt.Println("key : ", key,
				", value.LastName: ", *value.LastName,
				", value.Address[0].City : ", *value.Address[0].City,
				", value.PaymentMethod[0].PaymentToken", *value.PaymentMethod[0].PaymentToken)
		}

		c.JSON(200, users)
	})

	/**
	ex: 	http://localhost:5000/users-email/kylebanker2@gmail.com
	*/
	router.GET("/user/:email", func(c *gin.Context) {
		email := c.Param("email")
		user := service.GetUserByEmail(c, email)
		c.JSON(200, user)
	})
}
