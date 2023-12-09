package mongodb

import "go.mongodb.org/mongo-driver/mongo/options"

type IndexSpec struct {
	Key     map[string]interface{} `json:"key"`
	Options *options.IndexOptions  `json:"options"`
}
