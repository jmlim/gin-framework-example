package api

import (
	"fmt"
	"gin-framework-example/app/service"
	"gin-framework-example/websocket"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
)

func UserRouter(router *gin.Engine) {
	/**
	웹 소켓
	*/
	router.GET("/ws", func(c *gin.Context) {
		websocket.Wshandler(c.Writer, c.Request)
	})

	router.GET("/users", func(c *gin.Context) {
		users := service.GetUsers(c)

		// 조회
		for key, value := range users {
			fmt.Println("key : ", key,
				", value.LastName: ", *value.LastName,
				", value.Address[0].City : ", *value.Address[0].City,
				", value.PaymentMethod[0].PaymentToken", *value.PaymentMethod[0].PaymentToken)
		}

		c.JSON(http.StatusOK, users)
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

		c.JSON(http.StatusOK, users)
	})

	/**
	ex: 	http://localhost:5000/users-email/kylebanker2@gmail.com
	*/
	router.GET("/user/:email", func(c *gin.Context) {
		email := c.Param("email")
		user := service.GetUserByEmail(c, email)
		c.JSON(http.StatusOK, user)
	})

	/**
	ex: 	http://localhost:5000/users/search?lastName=kylebanker2@gmail.com
	*/
	router.GET("/users/search", func(c *gin.Context) {
		lastName := c.Query("lastName")
		firstName := c.Query("firstName")
		if lastName != "" {
			usersByLastName := service.GetUsersByLastName(c, lastName)
			c.JSON(http.StatusOK, usersByLastName)
		} else if firstName != "" {
			/*			usersByFirstNameStartsWith := service.GetUsersByFirstNameStartsWith(c, firstName)
						c.JSON(http.StatusOK, usersByFirstNameStartsWith)*/

			usersByFirstNameEndsWith := service.GetUsersByFirstNameEndsWith(c, firstName)
			c.JSON(http.StatusOK, usersByFirstNameEndsWith)
		}
	})

	router.GET("/users/paging-sample", func(c *gin.Context) {
		convertedPageInt, convertedLimitInt := getPageAndLimit(c)

		filter := bson.M{}
		limit := int64(convertedLimitInt)
		page := int64(convertedPageInt)
		userPaging := service.GetUsersPagingSample(c, page, limit, filter)

		c.JSON(http.StatusOK, userPaging)
	})
}

func getPageAndLimit(c *gin.Context) (convertedPageInt int, convertedLimitInt int) {
	queryPageValue := c.Query("page")
	if queryPageValue != "" {
		convertedPageInt, _ = strconv.Atoi(queryPageValue)
	}

	queryLimitValue := c.Query("limit")
	if queryLimitValue != "" {
		convertedLimitInt, _ = strconv.Atoi(queryLimitValue)
	}

	return convertedPageInt, convertedLimitInt
}
