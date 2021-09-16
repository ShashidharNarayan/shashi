package Repository

import (
	"github.com/ShashidharNarayan/bms/repository/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (*repo) Login(collection string, filter bson.M) (model.LoggedInObj, error) {
	result, err := Login(collection, filter)

	if err != nil {
		return model.LoggedInObj{}, err
	}

	return result, nil
}

func (*repo) GetBookings(collection string, filter bson.M) ([]model.LoggedInObj, error) {
	result, err := GetBookings(collection, filter)

	if err != nil {
		return nil, err
	}

	return result, nil
}
