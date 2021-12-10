package repo

import (
	"gin-framework-example/app/common"
	"gin-framework-example/app/domain"
	"gin-framework-example/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
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

func FindUsersPagingSample(c *gin.Context, page int64, limit int64, filter bson.M) common.PagingData {
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
		Data       interface{}           `json:"data"`
		Pagination common.PaginationData `json:"pagination"`
	}{
		Pagination: paginatedData.Pagination,
		Data:       users,
	}

	return payload
}

func FindUsersPagingSample2(c *gin.Context, request domain.PageRequest, filter bson.M) (*domain.UserPageData, error) {
	var users []*domain.User
	var usersCollection = getUsersCollection()

	findOptions := options.Find()
	findOptions.SetSkip((request.Page - 1) * request.Size)
	findOptions.SetLimit(request.Size)

	cursor, err := usersCollection.Find(c, filter, findOptions)
	defer cursor.Close(c)
	if err != nil {
		log.Printf("Find Error :  %v", err)
		return nil, err
	}

	err = cursor.All(c, &users)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}

	totalCount, err := usersCollection.CountDocuments(c, filter)
	if err != nil {
		log.Printf("CountDocuments Error: %v", err)
		return nil, err
	}
	page := getPage(totalCount, request)

	return &domain.UserPageData{users, page}, nil
}

func getPage(totalElement int64, request domain.PageRequest) *domain.Page {

	if totalElement == 0 {
		return nil
	}

	totalPages := int64(math.Ceil(float64(totalElement) / float64(request.Size)))
	var last = false

	if request.Page < totalPages {
		last = false
	} else {
		last = true
	}

	return &domain.Page{totalElement, totalPages, request.Size, last}
}
