package repo

import (
	"gin-framework-example/app/domain"
	"gin-framework-example/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func FindUserList(c *gin.Context) []domain.User {
	var usersCollection *mongo.Collection = db.OpenCollection(db.Client, "users")
	cursor, err := usersCollection.Find(c, bson.M{})
	defer cursor.Close(c)
	if err != nil {
		log.Fatal(err)
	}

	usersList := []domain.User{}
	for cursor.Next(c) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		usersList = append(usersList, user)
	}

	return usersList
}