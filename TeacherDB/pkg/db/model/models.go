package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Teacher struct {
	TeacherID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName  string             `json:"fullName" bson:"fullName"`
	Email     string             `json:"email" bson:"email"`
	Status    string             `json:"Status" bson:"Status"`
}
