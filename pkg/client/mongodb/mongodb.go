package mongodb

import (
	"clean_arhitecture_golang/internal/adapters/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClientDbImpl struct {
	db *mongo.Collection
}

func NewClient(ctx context.Context, host, port, username, password, dataBase, authDB string) (db.IDataBaseClient, error) {
	mongoDBUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s", username,password,host, port)
	isAuth := true
	if username == "" && password == "" {
		isAuth = false
		mongoDBUrl = fmt.Sprintf("mongodb://%s:%s", host, port)
	}

	uri := options.Client().ApplyURI(mongoDBUrl)
	if isAuth {
		if authDB == "" {
			authDB = dataBase
		}
		uri.SetAuth(options.Credential{
			AuthSource: authDB,
			Username: 	username,
			Password: 	password,
		})
	}

	connect, err := mongo.Connect(ctx, uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err)
	}

	if err := connect.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping to mongoDB due to error: %v", err)
	}

	return &mongoClientDbImpl{
			db: connect.Database(dataBase).Collection("blocks"),
		}, nil

}

func (m *mongoClientDbImpl) GetConnection() interface{} {
	return m.db
}


