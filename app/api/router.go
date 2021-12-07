package api

import (
	"fmt"
	"gin-framework-example/app/service"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/users", func(c *gin.Context) {
		userMapList := service.GetUserList(c)

		// 조회
		for key, value := range userMapList {
			fmt.Println("key : ", key,
				", value.LastName: ", *value.LastName,
				", value.Address[0].City : ", *value.Address[0].City,
				", value.PaymentMethod[0].PaymentToken", *value.PaymentMethod[0].PaymentToken)
		}

		c.JSON(200, userMapList)
	})
}
