package Repository

import (
	"github.com/ShashidharNarayan/bms/repository/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (*repo) GetTheatres(collection string, filter bson.M) ([]model.Theatre, error) {
	result, err := getTheatres(collection, filter)

	if err != nil {
		return nil, err
	}

	return result, nil
}
