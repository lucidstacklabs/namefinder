package apikey

import (
	"context"
	"fmt"
	"github.com/lucidstacklabs/namefinder/internal/pkg/secret"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Service struct {
	mongo *mongo.Collection
}

func NewService(mongo *mongo.Collection) *Service {
	return &Service{mongo: mongo}
}

func (s *Service) Create(ctx context.Context, request *CreationRequest, creatorID string) (*ApiKey, error) {
	nameExists, err := s.nameExists(ctx, request.Name)

	if err != nil {
		return nil, err
	}

	if nameExists {
		return nil, fmt.Errorf("api key %s already exists", request.Name)
	}

	apiKeySecret, err := secret.Generate(128)

	if err != nil {
		return nil, err
	}

	apiKey := &ApiKey{
		ID:        primitive.NewObjectID(),
		Name:      request.Name,
		Secret:    apiKeySecret,
		CreatorID: creatorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.mongo.InsertOne(ctx, apiKey)

	if err != nil {
		return nil, err
	}

	return apiKey, nil
}

func (s *Service) nameExists(ctx context.Context, name string) (bool, error) {
	count, err := s.mongo.CountDocuments(ctx, bson.M{"name": name})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
