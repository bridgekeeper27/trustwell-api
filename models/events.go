package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Content struct {
	GTIN           string `json:"gtin" binding:"required"`
	Lot            string `json:"lot" binding:"required"`
	BestByDate     string `json:"bestByDate" bson:"bestByDate"`
	ExpirationDate string `json:"expirationDate" bson:"expirationDate"`
}

type Event struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	CreatedBy string             `json:"createdBy" bson:"createdBy"`
	IsDeleted bool               `json:"isDeleted" bson:"isDeleted"`
	Type      string             `json:"type" binding:"required"`
	Contents  []Content          `json:"contents" binding:"required,dive,required"`
}
