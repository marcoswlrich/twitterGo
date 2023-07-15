package models

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	jwt.RegisteredClaims
}
