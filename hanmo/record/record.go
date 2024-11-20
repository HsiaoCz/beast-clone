package main

type Records struct {
	ID     string `bson:"_id" json:"id"`
	UserID string `bson:"user_id" json:"user_id"`
	BookID string `bson:"book_id" json:"book_id"`
}
