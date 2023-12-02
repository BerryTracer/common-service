package adapter

import "context"

type Cursor interface {
	All(ctx context.Context, results interface{}) error
	Close(ctx context.Context) error
	Decode(val interface{}) error
	Err() error
	Next(ctx context.Context) bool
	TryNext(ctx context.Context) bool
}
