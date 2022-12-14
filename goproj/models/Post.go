package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Username       string             `json:"username" bson:"username"`
	Title          string             `json:"title" `
	Text           string             `json:"text" `
	Answers        []string           `json:"answers" `
	AnswerUsername string             `json:"answerUsername"`
}
