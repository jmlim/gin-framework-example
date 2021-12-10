package service

import (
	"gin-framework-example/app/common"
	"gin-framework-example/app/domain"
	"gin-framework-example/app/repo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func GetUsersByLastName(c *gin.Context, lastname string) []domain.User {
	users := repo.FindUsersByLikeLastName(c, lastname)
	return users
}

func GetUsersByFirstNameStartsWith(c *gin.Context, firstname string) []domain.User {
	users := repo.FindUsersByFirstNameStartsWith(c, firstname)
	return users
}

func GetUsersByFirstNameEndsWith(c *gin.Context, firstname string) []domain.User {
	users := repo.FindUsersByFirstNameEndsWith(c, firstname)
	return users
}

func GetUsersPagingSample(c *gin.Context, page int64, limit int64, filter bson.M) common.PagingData {
	userPaging := repo.FindUsersPagingSample(c, page, limit, filter)
	return userPaging
}

/**
/*
	var episodes []bson.M
	if err = cursor.All(c, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)
*/
