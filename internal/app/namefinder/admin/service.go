package admin

import (
	"context"
	"errors"
	"fmt"
	"github.com/lucidstacklabs/namefinder/internal/pkg/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {
	mongo         *mongo.Collection
	authenticator *auth.Authenticator
}

func NewService(mongo *mongo.Collection, authenticator *auth.Authenticator) *Service {
	return &Service{mongo: mongo, authenticator: authenticator}
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

func (s *Service) GetToken(ctx context.Context, request *TokenRequest) (*TokenResponse, error) {
	result := s.mongo.FindOne(ctx, bson.M{"username": request.Username})

	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("invalid username and password combination")
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	var admin Admin

	if err := result.Decode(&admin); err != nil {
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(request.Password))

	if err != nil {
		return nil, fmt.Errorf("invalid username and password combination")
	}

	token, err := s.authenticator.GenerateAdminToken(admin.ID.Hex())

	if err != nil {
		return nil, err
	}

	return &TokenResponse{Token: token}, nil
}

func (s *Service) Get(ctx context.Context, adminID string) (*Admin, error) {
	id, err := primitive.ObjectIDFromHex(adminID)

	if err != nil {
		return nil, err
	}

	result := s.mongo.FindOne(ctx, bson.M{"_id": id})

	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("admin not found")
	}

	admin := &Admin{}

	if err := result.Decode(&admin); err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *Service) ChangePassword() {

}

func (s *Service) Add() {

}

func (s *Service) Delete() {

}

func (s *Service) List() {

}

func (s *Service) ResetPassword() {

}

func (s *Service) usernameExists(ctx context.Context, username string) (bool, error) {
	count, err := s.mongo.CountDocuments(ctx, bson.M{"username": username})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
