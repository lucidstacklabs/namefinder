package namespace

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Namespace struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	CreatorID string             `json:"creator_id" bson:"creator_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
