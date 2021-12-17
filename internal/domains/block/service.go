package block

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) *service {
	return  &service{storage: storage}
}

func (s *service) GetBlockById(id int) *Block {
	return s.storage.GetOne(id)
}

func (s *service) GetAll(ctx context.Context)  ([]Block, error) {
	return s.storage.GetAll(ctx)
}

func (s *service) AddNewBlock(ctx context.Context, b *DtoBlockRequest) (map[string]string, error) {
	lowLevelResponse := make(map[string]string)
	bl := &Block{
		ID:   primitive.NewObjectID(),
		Name: b.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Transactions: nil,
	}
	id, err := s.storage.Create(ctx, bl)
	if err != nil { return nil, err }
	lowLevelResponse["id"] = id

	return lowLevelResponse, nil
}

