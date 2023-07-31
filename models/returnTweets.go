package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnTweets struct {
	ID       primitive.ObjectID `bson:"_id"      json:"_id,omitempty"`
	UserID   string             `bson:"userid"   json:"userid,omitempty"`
	Mensagem string             `bson:"mensagem" json:"mensagem,omitempty"`
	Data     time.Time          `bson:"data"     json:"data,omitempty"`
}
