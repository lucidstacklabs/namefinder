package namespace

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Service struct {
	mongo *mongo.Collection
}

func NewService(mongo *mongo.Collection) *Service {
	return &Service{mongo: mongo}
}

func (s *Service) Create(ctx context.Context, request *CreationRequest, creatorID string) (*Namespace, error) {
	nameExists, err := s.nameExists(ctx, request.Name)

	if err != nil {
		return nil, err
	}

	if nameExists {
		return nil, fmt.Errorf("namespace %s already exists", request.Name)
	}

	namespace := &Namespace{
		ID:        primitive.NewObjectID(),
		Name:      request.Name,
		CreatorID: creatorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.mongo.InsertOne(ctx, namespace)

	if err != nil {
		return nil, err
	}

	return namespace, nil
}

func (s *Service) List(ctx context.Context, page int64, size int64) ([]*Namespace, error) {
	result, err := s.mongo.Find(ctx, bson.M{}, options.Find().SetSkip(page*size).SetLimit(size))

	if err != nil {
		return nil, err
	}

	namespaces := make([]*Namespace, 0)

	err = result.All(ctx, &namespaces)

	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (s *Service) Get(ctx context.Context, namespaceID string) (*Namespace, error) {
	id, err := primitive.ObjectIDFromHex(namespaceID)

	if err != nil {
		return nil, err
	}

	result := s.mongo.FindOne(ctx, bson.M{"_id": id})

	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("namespace not found")
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	var namespace Namespace

	err = result.Decode(&namespace)

	if err != nil {
		return nil, err
	}

	return &namespace, nil
}

func (s *Service) nameExists(ctx context.Context, name string) (bool, error) {
	count, err := s.mongo.CountDocuments(ctx, bson.M{"name": name})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
