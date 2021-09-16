package main

import (
	"io"
	"os"

	application "github.com/ShashidharNarayan/bms/application"
	"github.com/ShashidharNarayan/bms/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	app application.PostApplication = application.NewPostApplication()
)

func main() {
	router := gin.New()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(middlewares.CORSMiddleware())
	setupOutput()
	router.Use(gin.Recovery(), middlewares.Logger())

	router.POST("/api/register/", app.CreateUser)
	router.POST("/api/movie/", app.CreateMovie)
	router.POST("/api/insertshow/", app.CreateShow)
	router.POST("/api/inserttheatre/", app.CreateTheatre)
	router.POST("/api/addcoupons/", app.CreateCoupons)
	router.GET("/api/login/", app.Login)
	router.GET("/api/logout/", app.LogOut)
	router.GET("/api/movie/:pageNumber", app.GetMovie)
	router.GET("/api/movie/", app.GetMovie)
	router.GET("/api/checksession/", app.CheckSession)
	router.GET("/api/getbookings/:email", app.GetBookings) //middlewares.Auth(),
	router.GET("/api/getmoviecount/", app.GetMovieCount)
	router.GET("/api/getshow/:moviename", app.GetShows)
	router.GET("/api/getmovie/:moviename", app.GetMovieByName)
	router.GET("/api/getcoupons/", app.GetCoupons)
	router.GET("/api/gettheatres/", app.GetTheatres)
	router.GET("/api/getallmovies/", app.GetAllMovies)
	router.PUT("/api/movie/", app.UpdateMovie)
	router.PUT("/api/coupon/", app.UpdateCoupon)
	router.PUT("/api/updateshow/", app.UpdateShow)
	router.PUT("/api/updateuserbookings/", app.UpdateUser)
	router.PUT("/api/updateuserasadmin/", app.UpdateUser)
	router.DELETE("/api/movie/:id", app.DeleteMovie)
	router.DELETE("/api/deletecoupon/:couponCode", app.DeleteCoupons)
	router.Run("localhost:8080")
}

func setupOutput() {
	f, _ := os.Create("gin.log")
	io.MultiWriter(f, os.Stdout)
}
