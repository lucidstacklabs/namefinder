package namespace

import (
	"context"
	"fmt"
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder/apikey"
	apikey2 "github.com/lucidstacklabs/namefinder/internal/pkg/apikey"
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

func (s *ApiKeyAccessService) Delete(ctx context.Context, namespaceID string, request *ApiKeyAccessRequest) error {
	result, err := s.mongo.UpdateOne(ctx, bson.M{
		"namespace_id": namespaceID,
		"api_key_id":   request.ApiKeyID,
	}, bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
		"$pullAll": bson.M{
			"actions": request.Actions,
		},
	})

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("api key access to namespace does not exist")
	}

	return nil
}

func (s *ApiKeyAccessService) List(ctx context.Context, namespaceID string, page int64, size int64) ([]*ApiKeyAccessResponse, error) {
	result, err := s.mongo.Find(ctx, bson.M{
		"namespace_id": namespaceID,
	}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	accesses := make([]*ApiKeyAccess, 0)

	err = result.All(ctx, &accesses)

	if err != nil {
		return nil, err
	}

	if len(accesses) == 0 {
		return make([]*ApiKeyAccessResponse, 0), nil
	}

	apiKeyIDs := make([]string, len(accesses))

	for i, access := range accesses {
		apiKeyIDs[i] = access.ApiKeyID
	}

	apiKeys, err := s.apiKeyService.GetByIDs(ctx, apiKeyIDs)

	if err != nil {
		return nil, err
	}

	apiKeyMap := make(map[string]*apikey2.ApiKey, len(apiKeys))

	for _, apiKey := range apiKeys {
		apiKeyMap[apiKey.ID.Hex()] = apiKey
	}

	responses := make([]*ApiKeyAccessResponse, 0)

	for _, access := range accesses {
		responses = append(responses, &ApiKeyAccessResponse{
			Access: access,
			ApiKey: apiKeyMap[access.ApiKeyID],
		})
	}

	return responses, nil
}

func (s *ApiKeyAccessService) Destroy(ctx context.Context, namespaceID string, request *ApiKeyAccessDestroyRequest) error {
	result, err := s.mongo.DeleteOne(ctx, bson.M{
		"namespace_id": namespaceID,
		"api_key_id":   request.ApiKeyID,
	})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("api key access to namespace does not exist")
	}

	return nil
}
