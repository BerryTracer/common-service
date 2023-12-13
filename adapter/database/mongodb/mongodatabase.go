package mongodb

import (
	"context"

	"github.com/BerryTracer/common-service/adapter/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	clientOptions  *options.ClientOptions
	databaseName   string
	collectionName string
	collection     *mongo.Collection
}

func NewMongoDatabase(connStr, databaseStr, collectionStr string) *MongoDatabase {
	clientOptions := options.Client().ApplyURI(connStr)
	return &MongoDatabase{
		clientOptions:  clientOptions,
		databaseName:   databaseStr,
		collectionName: collectionStr,
	}
}

// Connect establishes the connection to the MongoDB server.
func (m *MongoDatabase) Connect(ctx context.Context) error {
	client, err := mongo.Connect(ctx, m.clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	m.collection = client.Database(m.databaseName).Collection(m.collectionName)
	return nil
}

// Disconnect closes the connection to the MongoDB server.
func (m *MongoDatabase) Disconnect(ctx context.Context) error {
	if m.collection != nil {
		return m.collection.Database().Client().Disconnect(ctx)
	}
	return nil
}

// GetCollection returns the MongoDB collection.
func (m *MongoDatabase) GetCollection() *mongo.Collection {
	return m.collection
}

func (m *MongoDatabase) CreateIndexes(ctx context.Context, indexSpecs []IndexSpec) error {
	for _, spec := range indexSpecs {
		indexModel := mongo.IndexModel{
			Keys:    spec.Key,
			Options: spec.Options,
		}
		if _, err := m.collection.Indexes().CreateOne(ctx, indexModel); err != nil {
			return err
		}
	}
	return nil
}

// Ensure that MongoDatabase implements the Database interface.
var _ database.Database = (*MongoDatabase)(nil)
