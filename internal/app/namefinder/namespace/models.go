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

type ApiKeyAccess struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	NamespaceID string             `json:"namespace_id" bson:"namespace_id"`
	ApiKeyID    string             `json:"api_key_id" bson:"api_key_id"`
	Actions     []string           `json:"actions" bson:"actions"`
	CreatorID   string             `json:"creator_id" bson:"creator_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type Action string

const (
	ActionCreate Action = "create"
	ActionDelete Action = "delete"
	ActionRead   Action = "read"
	ActionUpdate Action = "update"
)
