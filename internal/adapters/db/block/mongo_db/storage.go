package mongo_db

import (
	"clean_arhitecture_golang/internal/adapters/db"
	"clean_arhitecture_golang/internal/domains/block"
	"clean_arhitecture_golang/pkg/client/mongodb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type blockStorage struct {
	clientDb db.IDataBaseClient
}



func NewStorage() block.Storage {
	clientMongo, err := mongodb.NewClient(context.Background(),"172.20.0.17","27017","","","block-db","")
	if err != nil {
		panic(err)
	}

	return &blockStorage{
		clientDb: clientMongo,
	}
}

func (b blockStorage) GetOne(id int) *block.Block {
	panic("implement me")
}

func (b blockStorage) GetAll(ctx context.Context)  ([]block.Block, error) {
	client := b.clientDb.GetConnection().(*mongo.Collection)
	cursor, err := client.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("can't execute GET  query to mongo db: %v", err)
	}
	_ = cursor
	return nil, nil
}

func (b blockStorage) Create(ctx context.Context, block *block.Block) (string, error) {
	client := b.clientDb.GetConnection().(*mongo.Collection)
	one, err := client.InsertOne(ctx, &block)
	if err != nil { return "", fmt.Errorf("can't paste in mongo db: %v", err) }

	return one.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (b blockStorage) Delete(block *block.Block) {
	panic("implement me")
}



