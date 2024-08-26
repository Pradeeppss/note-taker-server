package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string   		  `json:"name" validate:"required,min=2,max=100"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	S3_path	   *string			  `json:"s3_path"`
	Subject_id *string			  `json:"subject_id"`
}