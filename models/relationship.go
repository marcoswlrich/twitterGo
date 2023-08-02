package models

type Relationship struct {
	UserID         string `bson:"usuarioid"      json:"usuarioid"`
	UserRelationID string `bson:"userrelationid" json:"userrelationid"`
}
