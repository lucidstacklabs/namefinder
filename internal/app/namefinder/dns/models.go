package dns

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Record struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Type        string             `bson:"type" json:"type"`
	Value       string             `bson:"value" json:"value"`
	TTL         uint32             `bson:"ttl" json:"ttl"`
	Class       string             `bson:"class" json:"class"`
	CreatorType ActorType          `bson:"actor_type" json:"actor_type"`
	CreatorID   string             `json:"creator_id" bson:"creator_id" json:"creator_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type ActorType string

const (
	ActorTypeAdmin  ActorType = "admin"
	ActorTypeApiKey ActorType = "apikey"
)
