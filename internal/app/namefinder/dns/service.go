package dns

import (
	"context"
	"fmt"
	"github.com/lucidstacklabs/namefinder/internal/app/namefinder/namespace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s *RecordService) List(ctx context.Context, namespaceID string, page int64, size int64) ([]*Record, error) {
	result, err := s.mongo.Find(ctx, bson.M{
		"namespace_id": namespaceID,
	}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	records := make([]*Record, 0)

	err = result.All(ctx, &records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *RecordService) Get() {

}

func (s *RecordService) Update() {

}

func (s *RecordService) Delete() {

}
