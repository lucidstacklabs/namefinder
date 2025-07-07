package namespace

import (
	"context"
	"fmt"
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder/apikey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ApiKeyAccessService struct {
	mongo         *mongo.Collection
	service       *Service
	apiKeyService *apikey.Service
}

func NewApiKeyAccessService(mongo *mongo.Collection, service *Service, apiKeyService *apikey.Service) *ApiKeyAccessService {
	return &ApiKeyAccessService{mongo: mongo, service: service, apiKeyService: apiKeyService}
}

func (s *ApiKeyAccessService) Add(ctx context.Context, namespaceID string, request *ApiKeyAccessRequest, creatorID string) error {
	namespaceExists, err := s.service.Exists(ctx, namespaceID)

	if err != nil {
		return err
	}

	if !namespaceExists {
		return fmt.Errorf("namespace not found")
	}

	apiKeyExists, err := s.apiKeyService.Exists(ctx, request.ApiKeyID)

	if err != nil {
		return err
	}

	if !apiKeyExists {
		return fmt.Errorf("api key not found")
	}

	_, err = s.mongo.UpdateOne(ctx, bson.M{
		"namespace_id": namespaceID,
		"api_key_id":   request.ApiKeyID,
	}, bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
		"$setOnInsert": bson.M{
			"namespace_id": namespaceID,
			"api_key_id":   request.ApiKeyID,
			"creator_id":   creatorID,
			"created_at":   time.Now(),
		},
		"$addToSet": bson.M{
			"actions": bson.M{
				"$each": request.Actions,
			},
		},
	}, options.Update().SetUpsert(true))

	return err
}
