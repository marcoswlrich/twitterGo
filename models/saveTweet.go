package models

import "time"

type SaveTweet struct {
	UserID   string    `bson:"userid"   json:"userid,omitempty"`
	Mensagem string    `bson:"mensagem" json:"mensagem,omitempty"`
	Data     time.Time `bson:"data"     json:"data,omitempty"`
}
