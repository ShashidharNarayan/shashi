package Repository

import (
	"github.com/ShashidharNarayan/bms/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Login(collection string, filter bson.M) (model.LoggedInObj, error)
	GetMovies(collection string, query bson.M, options *options.FindOptions) ([]model.Movie, error)
	Create(collection string, data interface{}, filter bson.M) (interface{}, error)
	Update(collection string, data bson.M, query bson.M) (*mongo.UpdateResult, error)
	Delete(collection string, filter bson.M) (*mongo.DeleteResult, error)
	GetBookings(collection string, filter bson.M) ([]model.LoggedInObj, error)
	GetMovieCount(collection string) (int64, error)
	GetShowsOfTheMovie(collection string, movieName string) ([]model.MovieShow, error)
	GetMovieByName(collection string, filter bson.M) (model.Movie, error)
	GetMovieCoupons(collection string) ([]model.Coupon, error)
	GetTheatres(collection string, filter bson.M) ([]model.Theatre, error)
}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}
