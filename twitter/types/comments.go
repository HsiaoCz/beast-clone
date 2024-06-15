package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// seem twitter no comments
// or comments are posts
type Comment struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID   primitive.ObjectID `bson:"userID" json:"userID"`
	PostID   primitive.ObjectID `bson:"postID" json:"postID"`
	ParentID primitive.ObjectID `bson:"parentID" json:"parentID"`
	Content  string             `bson:"content" json:"content"`
	CreateAt time.Time          `bson:"createAt" json:"createAt"`
	Watches  string             `bson:"watches" json:"watches"`
	Likes    string             `bson:"likes" json:"likes"`
	CommetCt string             `bson:"commentCt" json:"commentCt"`
}

type CreateCommentParams struct {
	UserID   primitive.ObjectID `json:"userID"`
	PostID   primitive.ObjectID `json:"postID"`
	ParentID primitive.ObjectID `json:"parentID,omitempty"`
	Content  string             `json:"content"`
}

func NewCommentsFromParams(param CreateCommentParams) *Comment {
	return &Comment{
		UserID:   param.UserID,
		PostID:   param.PostID,
		ParentID: param.ParentID,
		Content:  param.Content,
		CreateAt: time.Now(),
		Watches:  "0",
		Likes:    "0",
		CommetCt: "0",
	}
}
