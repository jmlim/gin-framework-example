package service

import (
	"gin-framework-example/app/domain"
	"gin-framework-example/app/repo"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) []domain.User {
	users := repo.FindUsers(c)
	return users
}

func GetUsersByName(c *gin.Context, name string) []domain.User {
	users := repo.FindUsersByName(c, name)
	return users
}

func GetUserByEmail(c *gin.Context, email string) domain.User {
	user := repo.FindUserByEmail(c, email)
	return user
}

/**
/*
	var episodes []bson.M
	if err = cursor.All(c, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)
*/
