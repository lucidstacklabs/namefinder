package admin

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	mongo *mongo.Collection
}

func NewService(mongo *mongo.Collection) *Service {
	return &Service{mongo: mongo}
}

func (s *Service) Init(ctx context.Context) {

}

func (s *Service) GetToken() {

}
