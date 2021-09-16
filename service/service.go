package Service

import (
	Repository "github.com/ShashidharNarayan/bms/repository"
	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct{}

type Service interface {
	ValidateMovie(data *model.Movie) error
	ValidateTheatre(data *model.Theatre) error
	ValidateCoupon(data *model.Coupon) error
	ValidateShow(data *model.Show) error
	ValidateUser(data *model.User) error
	ValidateLogin(data *model.User) error

	CreateMovie(data model.Movie) (interface{}, error)
	UpdateMovie(movie model.Movie) (interface{}, error)
	DeleteMovie(id string) (*mongo.DeleteResult, error)
	GetMovies(pageNumber int64) ([]model.Movie, error)
	GetAllMovies() ([]model.Movie, error)

	CreateShow(data model.Show) (interface{}, error)
	UpdateShow(movie model.Show) (interface{}, error)
	DeleteShow(id string) (*mongo.DeleteResult, error)
	GetShowsOfTheMovie(movieName string) ([]model.MovieShow, error)

	CreateTheatre(data model.Theatre) (interface{}, error)
	UpdateTheatre(movie model.Theatre) (interface{}, error)
	DeleteTheatre(id string) (*mongo.DeleteResult, error)
	GetTheatres() ([]model.Theatre, error)

	CreateCoupon(data model.Coupon) (interface{}, error)
	UpdateCoupon(movie model.Coupon) (interface{}, error)
	DeleteCoupon(id string) (*mongo.DeleteResult, error)
	GetCoupon() ([]model.Coupon, error)

	CreateUser(data model.User, c *gin.Context) (interface{}, error)
	UpdateUser(movie model.User) (interface{}, error)
	DeleteUser(id string) (*mongo.DeleteResult, error)

	Login(email string, password string, c *gin.Context) (model.LoggedInObj, error)
	GetBookings(email string) ([]model.LoggedInObj, error)
	GetMovieCount() (int64, error)
	GetMovieByName(movieName string) (model.Movie, error)
}

var (
	repo Repository.Repository = Repository.NewRepository()
)

func NewService() Service {
	return &service{}
}
