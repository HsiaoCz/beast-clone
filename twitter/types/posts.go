package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"userID" json:"userID"`
	Content    string             `bson:"content" json:"content"`
	CreateAt   time.Time          `bson:"createAt" json:"createAt"`
	StaticPath []string           `bson:"staticPath,omitempty" json:"staticPath,omitempty"`
	Likes      string             `bson:"likes" json:"likes"`
	Comments   string             `bson:"comments" json:"comments"`
	Watches    string             `bson:"watches" json:"watches"`
}

type CreatePostParams struct {
	Content    string   `json:"content"`
	StaticPath []string `json:"staticPath,omitempty"`
}

func NewPostFromParams(param CreatePostParams) *Post {
	return &Post{
		Content:    param.Content,
		StaticPath: param.StaticPath,
		CreateAt:   time.Now(),
		Likes:      "0",
		Comments:   "0",
		Watches:    "0",
	}
}
