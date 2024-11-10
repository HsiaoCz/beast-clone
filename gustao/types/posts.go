package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Posts struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	PostID string             `bson:"post_id" json:"post_id"`
	
}
