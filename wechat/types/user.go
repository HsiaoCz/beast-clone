package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
}

type UserInfo struct {
	ID      primitive.ObjectID
	Email   string
	IsAdmin bool
}
