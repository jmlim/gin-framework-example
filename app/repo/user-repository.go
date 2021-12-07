package repo

import (
	"gin-framework-example/app/domain"
	"gin-framework-example/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func getUsersCollection() *mongo.Collection {
	var usersCollection *mongo.Collection = db.OpenCollection(db.Client, "users")
	return usersCollection
}

func FindUsers(c *gin.Context) []domain.User {
	var usersCollection = getUsersCollection()
	cursor, err := usersCollection.Find(c, bson.M{})
	defer cursor.Close(c)
	if err != nil {
		log.Fatal(err)
	}

	users := []domain.User{}
	for cursor.Next(c) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}

func FindUsersByName(c *gin.Context, name string) []domain.User {
	var usersCollection = getUsersCollection()
	cursor, err := usersCollection.Find(c, bson.M{"username": name})
	defer cursor.Close(c)
	if err != nil {
		log.Fatal(err)
	}

	usersByName := []domain.User{}
	for cursor.Next(c) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		usersByName = append(usersByName, user)
	}

	return usersByName
}

func FindUserByEmail(c *gin.Context, email string) domain.User {
	var usersCollection = getUsersCollection()
	var user domain.User
	err := usersCollection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}
