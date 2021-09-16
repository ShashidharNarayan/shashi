package Repository

import (
	"github.com/ShashidharNarayan/bms/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (*repo) GetMovies(collection string, query bson.M, options *options.FindOptions) ([]model.Movie, error) {
	result, err := getMovies(collection, query, options)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*repo) GetMovieCount(collection string) (int64, error) {
	result, err := GetMovieCount(collection)

	if err != nil {
		return 0, err
	}

	return result, nil
}

func (*repo) GetMovieByName(collection string, filter bson.M) (model.Movie, error) {
	result, err := GetMovie(collection, filter)

	if err != nil {
		return model.Movie{}, err
	}

	return result, nil
}
