package repo

import (
	"gin-framework-example/app/common"
	"gin-framework-example/app/domain"
	"gin-framework-example/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func FindUserByEmail(c *gin.Context, email string) domain.User {
	var usersCollection = getUsersCollection()
	var user domain.User
	err := usersCollection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func FindUsersByLikeLastName(c *gin.Context, lastName string) []domain.User {
	var usersCollection = getUsersCollection()
	cursor, err := usersCollection.Find(c, bson.M{"lastName": primitive.Regex{Pattern: lastName, Options: ""}})
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

func FindUsersByFirstNameStartsWith(c *gin.Context, firstName string) []domain.User {
	var usersCollection = getUsersCollection()
	cursor, err := usersCollection.Find(c, bson.M{"firstName": primitive.Regex{Pattern: "^" + firstName, Options: "i"}})
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

func FindUsersByFirstNameEndsWith(c *gin.Context, firstName string) []domain.User {
	var usersCollection = getUsersCollection()
	cursor, err := usersCollection.Find(c, bson.M{"firstName": primitive.Regex{Pattern: firstName + "$", Options: "i"}})
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

func FindUsersPagingSample(c *gin.Context, page int64, limit int64, filter bson.M) domain.UserPaging {
	var usersCollection = getUsersCollection()
	var users []domain.User

	projection := bson.D{
		{"username", 1},
		{"email", 1},
		{"firstName", 1},
		{"lastName", 1},
		{"hashPassword", 1},
		{"address", 1},
		{"paymentMethod", 1},
	}

	paginatedData, err := common.New(usersCollection).Context(c).Limit(limit).Page(page).Sort("username", -1).Select(projection).Filter(filter).Decode(&users).Find()
	if err != nil {
		log.Fatal(err)
	}

	var payload = struct {
		Data       []domain.User         `json:"data"`
		Pagination common.PaginationData `json:"pagination"`
	}{
		Pagination: paginatedData.Pagination,
		Data:       users,
	}

	return payload
}
