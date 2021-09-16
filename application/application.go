package application

import (
	Service "github.com/ShashidharNarayan/bms/service"
	"github.com/gin-gonic/gin"
)

type application struct{}

type PostApplication interface {
	CreateCoupons(c *gin.Context)
	DeleteCoupons(c *gin.Context)
	UpdateCoupon(c *gin.Context)
	GetCoupons(c *gin.Context)

	GetAllMovies(c *gin.Context)
	GetMovieByName(c *gin.Context)
	GetMovieCount(c *gin.Context)
	DeleteMovie(c *gin.Context)
	UpdateMovie(c *gin.Context)
	GetMovie(c *gin.Context)
	CreateMovie(c *gin.Context)

	DeleteShow(c *gin.Context)
	UpdateShow(c *gin.Context)
	CreateShow(c *gin.Context)
	GetShows(c *gin.Context)

	DeleteTheatre(c *gin.Context)
	UpdateTheatre(c *gin.Context)
	GetTheatres(c *gin.Context)
	CreateTheatre(c *gin.Context)

	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetBookings(c *gin.Context)
	LogOut(c *gin.Context)
	CheckSession(c *gin.Context)
	Login(c *gin.Context)
	CreateUser(c *gin.Context)
}

var (
	app Service.Service = Service.NewService()
)

func NewPostApplication() PostApplication {
	return &application{}
}
