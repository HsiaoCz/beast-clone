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
	ParentID primitive.ObjectID `bson:"parentID,omitempty" json:"parentID,omitempty"`
	Content  string             `bson:"content" json:"content"`
	CreateAt time.Time          `bson:"createAt" json:"createAt"`
	Watches  string             `bson:"watches" json:"watches"`
	Likes    string             `bson:"likes" json:"likes"`
	CommetCt string             `bson:"commentCt" json:"commentCt"`
}

type CreateCommentParams struct {
	PostID   string `json:"postID"`
	ParentID string `json:"parentID,omitempty"`
	Content  string `json:"content"`
}

func (param CreateCommentParams) Validate() map[string]string {
	errors := map[string]string{}
	if err := isValidId(param.PostID); err != nil {
		errors["postID"] = "invalid post identity"
	}
	if param.ParentID != "" {
		if err := isValidId(param.ParentID); err != nil {
			errors["parentID"] = "invalid parent identity"
		}
	}
	return errors
}

func isValidId(id string) error {
	_, err := primitive.ObjectIDFromHex(id)
	return err
}

func NewCommentsFromParams(param CreateCommentParams, userID primitive.ObjectID) *Comment {
	comment := &Comment{
		UserID:   userID,
		Content:  param.Content,
		CreateAt: time.Now(),
		Watches:  "0",
		Likes:    "0",
		CommetCt: "0",
	}
	postID, _ := primitive.ObjectIDFromHex(param.PostID)
	comment.PostID = postID

	if param.ParentID != "" {
		parentID, _ := primitive.ObjectIDFromHex(param.ParentID)
		comment.ParentID = parentID
	}
	return comment
}
