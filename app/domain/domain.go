package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product struct
type Product struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name     string             `json:"name,omitempty" bson:"name"`
	Quantity float64            `json:"quantity,omitempty" bson:"quantity"`
	Price    float64            `json:"price,omitempty" bson:"price"`
}

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Username      *string            `json:"username" validate:"required,min=2,max=50"`
	Email         *string            `json:"email" validate:"required,min=2,max=100"`
	FirstName     *string            `json:"firstName" validate:"required"`
	LastName      *string            `json:"lastName" validate:"required"`
	HashPassword  *string            `json:"hashPassword" validate:"required"`
	Address       []Address          `json:"address"`
	PaymentMethod []PaymentMethod    `json:"paymentMethod"`
}

type Address struct {
	Name   *string `json:"name"`
	Street *string `json:"street"`
	City   *string `json:"city"`
	Zip    int32   `json:"zip"`
}

type PaymentMethod struct {
	Name         *string `json:"name"`
	PaymentToken *string `json:"paymentToken"`
}
