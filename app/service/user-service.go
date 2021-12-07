package service

import (
	"fmt"
	"gin-framework-example/app/domain"
	"gin-framework-example/app/repo"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) []domain.User {
	usersList := repo.FindUserList(c)
	for key, value := range usersList {
		fmt.Println("key : ", key,
			", value.LastName: ", *value.LastName,
			", value.Address[0].City : ", *value.Address[0].City,
			", value.PaymentMethod[0].PaymentToken", *value.PaymentMethod[0].PaymentToken)
	}
	return usersList
}

/**
/*
	var episodes []bson.M
	if err = cursor.All(c, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)
*/
