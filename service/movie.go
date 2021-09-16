package Service

import (
	"github.com/pkg/errors"

	"github.com/ShashidharNarayan/bms/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection string = "movie"

func (*service) ValidateMovie(data *model.Movie) error {
	if data == nil {
		err := errors.New("Movie Data Cannot Be Empty")
		return err
	}

	if data.MovieName == "" {
		err := errors.New("Movie Name Cannot Be Empty")
		return err
	}

	if data.MovieDetail == "" {
		err := errors.New("Movie Detail Cannot Be Empty")
		return err
	}

	if data.MovieImageUrl == "" {
		err := errors.New("Movie Image Url Cannot Be Empty")
		return err
	}

	if data.MovieTrailer == "" {
		err := errors.New("Movie Trailer Cannot Be Empty")
		return err
	}

	if data.ScreenedAt.IsZero() {
		err := errors.New("Movie Trailer Cannot Be Empty")
		return err
	}

	return nil
}

func (*service) CreateMovie(data model.Movie) (interface{}, error) {
	data.ID = primitive.NewObjectID()
	filter := bson.M{"movie_name": data.MovieName}
	result, err := repo.Create(collection, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) GetMovies(pageNumber int64) ([]model.Movie, error) {
	opts := options.Find()
	opts.SetSort(bson.M{"screened_at": 1})
	opts.SetLimit(5)
	opts.SetSkip((pageNumber - 1) * 5)
	filter := bson.M{}
	result, err := repo.GetMovies(collection, filter, opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) GetAllMovies() ([]model.Movie, error) {
	opts := options.Find()
	opts.SetSort(bson.M{"screened_at": 1})
	filter := bson.M{}
	result, err := repo.GetMovies(collection, filter, opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) UpdateMovie(movie model.Movie) (interface{}, error) {
	data := bson.M{
		"$set": bson.M{
			"movie_detail":   movie.MovieDetail,
			"movie_name":     movie.MovieName,
			"movie_trailer":  movie.MovieTrailer,
			"movie_imageurl": movie.MovieImageUrl,
			"screened_at":    movie.ScreenedAt,
		},
	}

	filter := bson.M{"movie_name": movie.MovieName}
	result, err := repo.Update(collection, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) DeleteMovie(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	result, err := repo.Delete(collection, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) GetMovieCount() (int64, error) {
	result, err := repo.GetMovieCount(collection)

	if err != nil {
		return 0, err
	}

	return result, nil
}

func (*service) GetMovieByName(movieName string) (model.Movie, error) {
	filter := bson.M{"movie_name": movieName}
	result, err := repo.GetMovieByName(collection, filter)

	if err != nil {
		return model.Movie{}, err
	}

	return result, nil
}
