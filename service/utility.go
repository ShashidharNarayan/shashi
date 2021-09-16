package Service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (*service) Create(collection string, data interface{}, filter bson.M) (interface{}, error) {
	result, err := repo.Create(collection, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) Update(collection string, data bson.M, query bson.M) (*mongo.UpdateResult, error) {
	result, err := repo.Update(collection, data, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) Delete(collection string, filter bson.M) (*mongo.DeleteResult, error) {
	result, err := repo.Delete(collection, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
