package Repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (*repo) Create(collection string, data interface{}, filter bson.M) (interface{}, error) {
	result, err := createDocument(collection, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*repo) Update(collection string, data bson.M, query bson.M) (*mongo.UpdateResult, error) {
	result, err := updateDocument(collection, data, query)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*repo) Delete(collection string, filter bson.M) (*mongo.DeleteResult, error) {
	res, err := deleteDocument(collection, filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}
