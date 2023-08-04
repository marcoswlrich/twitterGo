package models

type Relationship struct {
	UserID         string `bson:"userid"         json:"userid"`
	UserRelationID string `bson:"userrelationid" json:"userrelationid"`
}
