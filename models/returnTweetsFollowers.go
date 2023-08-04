package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnTweetsFollowers struct {
	ID             primitive.ObjectID `bson:"_id"            json:"_id,omitempty"`
	UserID         string             `bson:"userid"         json:"userid,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userrelationid,omitempty"`
	Tweet          struct {
		Mensagem string    `bson:"mensagem" json:"mensagem,omitempty"`
		Data     time.Time `bson:"data" json:"data,omitempty"`
		ID       string    `bson:"_id" json:"_id,omitempty"`
	}
}
