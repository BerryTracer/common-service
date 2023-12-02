package adapter

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAdapter interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error)
}

type MongoAdapterImpl struct {
	collection *mongo.Collection
}

func NewMongoAdapter(collection *mongo.Collection) *MongoAdapterImpl {
	return &MongoAdapterImpl{collection: collection}
}

func (m *MongoAdapterImpl) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return m.collection.InsertOne(ctx, document, opts...)
}

func (m *MongoAdapterImpl) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.collection.UpdateOne(ctx, filter, update, opts...)
}

func (m *MongoAdapterImpl) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.collection.DeleteOne(ctx, filter, opts...)
}

func (m *MongoAdapterImpl) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return m.collection.FindOne(ctx, filter, opts...)
}

func (m *MongoAdapterImpl) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	return m.collection.Find(ctx, filter, opts...)
}

// Ensure that MongoAdapterImpl implements MongoAdapter
var _ MongoAdapter = (*MongoAdapterImpl)(nil)
