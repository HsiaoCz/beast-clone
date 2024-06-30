package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BookName string             `bson:"bookName"`
	Content  string             `bson:"content" json:"content"`
	Auther   string             `bson:"auther"`
	Avatar   string             `bson:"avatar" json:"avatar"`
	Path     string             `bson:"path" json:"path"`
}
