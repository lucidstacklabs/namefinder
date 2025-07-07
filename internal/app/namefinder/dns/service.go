package dns

import (
	"context"
	"fmt"
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder/namespace"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RecordService struct {
	mongo            *mongo.Collection
	namespaceService *namespace.Service
}

func NewRecordService(mongo *mongo.Collection, namespaceService *namespace.Service) *RecordService {
	return &RecordService{mongo: mongo, namespaceService: namespaceService}
}

func (s *RecordService) Add(ctx context.Context, namespaceID string, request *RecordAdditionRequest, creatorType ActorType, creatorID string) (*Record, error) {
	namespaceExists, err := s.namespaceService.Exists(ctx, namespaceID)

	if err != nil {
		return nil, err
	}

	if !namespaceExists {
		return nil, fmt.Errorf("namespace not found")
	}

	record := &Record{
		ID:          primitive.NewObjectID(),
		NamespaceID: namespaceID,
		Name:        request.Name,
		Type:        request.Type,
		Value:       request.Value,
		TTL:         request.TTL,
		Class:       request.Class,
		CreatorType: creatorType,
		CreatorID:   creatorID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = s.mongo.InsertOne(ctx, request)

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (s *RecordService) List() {

}

func (s *RecordService) Get() {

}

func (s *RecordService) Update() {

}

func (s *RecordService) Delete() {

}
