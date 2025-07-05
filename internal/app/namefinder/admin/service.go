package admin

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {
	mongo *mongo.Collection
}

func NewService(mongo *mongo.Collection) *Service {
	return &Service{mongo: mongo}
}

func (s *Service) Init(ctx context.Context, request *InitRequest) (*Admin, error) {
	count, err := s.mongo.CountDocuments(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	if count != 0 {
		return nil, fmt.Errorf("you are not allowed to perform this action")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	admin := &Admin{
		ID:        primitive.NewObjectID(),
		Username:  request.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.mongo.InsertOne(ctx, admin)

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *Service) GetToken() {

}
