package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"` //json:"_"
	Title   string             `bson:"title" json:"title"`
	Content string             `bson:"content" json:"content"`
	UserID  primitive.ObjectID `bson:"userId" json:"-"` //json:"_"
}
type PostDTO struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	UserID  string   `json:"userId"`
	User    *UserDTO `json:"user,omitempty"` // populated optionally
}
