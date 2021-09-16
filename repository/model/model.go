package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	MovieName     string             `bson:"movie_name" json:"movie_name"`
	MovieTrailer  string             `bson:"movie_trailer" json:"movie_trailer"`
	MovieDetail   string             `bson:"movie_detail" json:"movie_detail"`
	MovieImageUrl string             `bson:"movie_imageurl" json:"movie_imageurl"`
	ScreenedAt    time.Time          `bson:"screened_at" json:"screened_at"`
}

type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Name       string             `bson:"name" json:"name"`
	Gender     string             `bson:"gender" json:"gender"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password"`
	IsAdmin    bool               `bson:"is_admin" json:"is_admin"`
	MyBookings Bookings           `bson:"my_bookings" json:"my_bookings"`
}

type Show struct {
	ID                  primitive.ObjectID `bson:"_id" json:"_id"`
	TheatreName         string             `bson:"theatre_name" json:"theatre_name"`
	MovieName           string             `bson:"movie_name" json:"movie_name"`
	ShowDate            time.Time          `bson:"show_date" json:"show_date"`
	StartTime           Start_time         `bson:"start_time" json:"start_time"`
	SilverCategorySeats int                `bson:"silver_category_seats" json:"silver_category_seats"`
	GoldCategorySeats   int                `bson:"gold_category_seats" json:"gold_category_seats"`
	SilverCategoryPrice int                `bson:"silver_category_price" json:"silver_category_price"`
	GoldCategoryPrice   int                `bson:"gold_category_price" json:"gold_category_price"`
	MovieStartTime      string             `bson:"movie_start_time" json:"movie_start_time"`
	Seats               [][]Ticket         `bson:"seats" json:"seats"`
	SeatLimit           int                `bson:"seatlimit" json:"seatlimit"`
}

type Start_time struct {
	Hour   int `bson:"hour" json:"hour"`
	Minute int `bson:"minute" json:"minute"`
}

type Ticket struct {
	IsBooked bool   `bson:"is_booked" json:"is_booked"`
	Email    string `bson:"email" json:"email"`
	Price    int    `bson:"price" json:"price"`
	Category string `bson:"category" json:"category"`
}

type Theatre struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	TheatreName string             `bson:"theatre_name" json:"theatre_name"`
	TheatreUrl  string             `bson:"theatre_url" json:"theatre_url"`
}

type Bookings struct {
	MovieTime string    `bson:"movie_time" json:"movie_time"`
	MovieName string    `bson:"movie_name" json:"movie_name"`
	Date      time.Time `bson:"date" json:"date"`
	Seats     []int     `bson:"seats" json:"seats"`
}

type Coupon struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	CouponCode  string
	CouponPrice int
}

type MovieShow struct {
	ID                  primitive.ObjectID `bson:"_id" json:"_id"`
	TheatreName         string             `bson:"theatre_name" json:"theatre_name"`
	MovieName           string             `bson:"movie_name" json:"movie_name"`
	ShowDate            time.Time          `bson:"show_date" json:"show_date"`
	StartTime           Start_time         `bson:"start_time" json:"start_time"`
	SilverCategorySeats int                `bson:"silver_category_seats" json:"silver_category_seats"`
	GoldCategorySeats   int                `bson:"gold_category_seats" json:"gold_category_seats"`
	SilverCategoryPrice int                `bson:"silver_category_price" json:"silver_category_price"`
	GoldCategoryPrice   int                `bson:"gold_category_price" json:"gold_category_price"`
	MovieStartTime      string             `bson:"movie_start_time" json:"movie_start_time"`
	Seats               [][]Ticket         `bson:"seats" json:"seats"`
	SeatLimit           int                `bson:"seatlimit" json:"seatlimit"`
	MovieObj
	TheatreObj
}

type MovieObj struct {
	MovieName     string    `bson:"movie_name" json:"movie_name" binding:"required"`
	MovieTrailer  string    `bson:"movie_trailer" json:"movie_trailer" binding:"required"`
	MovieDetail   string    `bson:"movie_detail" json:"movie_detail" binding:"required"`
	MovieImageUrl string    `bson:"movie_imageurl" json:"movie_imageurl" binding:"required"`
	ScreenedAt    time.Time `bson:"screened_at" json:"screened_at" binding:"required"`
}

type LoggedInObj struct {
	Name       string   `bson:"name" json:"name"`
	Gender     string   `bson:"gender" json:"gender"`
	Email      string   `bson:"email" json:"email"`
	IsAdmin    bool     `bson:"is_admin" json:"is_admin"`
	MyBookings Bookings `bson:"my_bookings" json:"my_bookings"`
}

type TheatreObj struct {
	TheatreName string `bson:"theatre_name" json:"theatre_name"`
	TheatreUrl  string `bson:"theatre_url" json:"theatre_url"`
}
