package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email" json:"email"`
	Content  string             `bson:"content" json:"content"`
	Avatar   string             `bson:"avatar" json:"avatar"`
	ReadTime string             `bson:"readTime" json:"readTime"`
	IsAdmin  bool               `bson:"isAdmin" json:"isAdmin"`
}

//