package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostBody struct {
	CustomerId string `json:"account_id"`
	Amount     string `json:"amount"`
}

type DTO struct {
	PaymentId string `json:"id"`
	Amount    string `json:"amount"`
	Status    string `json:"status"`
}

type DAO struct {
	ID         primitive.ObjectID         `bson:"_id,omitempty"`
	Amount     int64                      `bson:"amount"`
	CustomerId string                     `bson:"customerId"`
	PaymentId  string                     `bson:"paymentId"`
}
