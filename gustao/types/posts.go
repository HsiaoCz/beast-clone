package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	PostID      string             `bson:"post_id" json:"post_id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	Title       string             `bson:"title" json:"title"`
	Brief       string             `bson:"brief" json:"brief"`
	Auther      string             `bson:"auther" json:"auther"`
	Likes       string             `bson:"likes" json:"likes"`
	Views       string             `bson:"views" json:"views"`
	Content     string             `bson:"content" json:"content"`
	Tags        []string           `bson:"tags" json:"tags"`
	Featured    string             `bson:"featured" json:"featured"`
	AutherBio   string             `bson:"auther_bio" json:"auther_bio"`
	Attachments []string           `bson:"attachments" json:"attachments"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"create_at" json:"create_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
