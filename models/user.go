package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"  json:"id"`
	Name           string             `bson:"name"           json:"name,omitempty"`
	Apelidos       string             `bson:"apelidos"       json:"apelidos,omitempty"`
	DataNascimento time.Time          `bson:"dataNascimento" json:"dataNascimento,omitempty"`
	Email          string             `bson:"email"          json:"email"`
	Password       string             `bson:"password"       json:"password,omitempty"`
	Avatar         string             `bson:"avatar"         json:"avatar,omitempty"`
	Banner         string             `bson:"banner"         json:"banner,omitempty"`
	Biografia      string             `bson:"biografia"      json:"biografia,omitempty"`
	Publication    string             `bson:"publication"    json:"publication,omitempty"`
	SiteWeb        string             `bson:"siteweb"        json:"siteweb,omitempty"`
}
