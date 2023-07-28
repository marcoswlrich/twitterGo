package models

type Tweet struct {
	Mensagem string `bson:"mensagem" json:"mensage"`
}
