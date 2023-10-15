package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	FileID string             `bson:"fileId"`
	Hash    string             `bson:"hash"`
}
