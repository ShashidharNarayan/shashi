package application

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/gin-gonic/gin"
)

func (*application) CreateMovie(c *gin.Context) {
	var movie model.Movie
	movie, error := decodeData(c)
	if error != nil {
		c.JSON(http.StatusInternalServerError, error)
	}

	err := app.ValidateMovie(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	result, err := app.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) GetMovie(c *gin.Context) {
	page := c.Param("pageNumber")
	if page != "" {
		pageNumber, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		result, err := app.GetMovies(pageNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusCreated, result)
	} else {
		c.JSON(http.StatusCreated, "Page Number Cannot be Empty")
	}
}

func (*application) UpdateMovie(c *gin.Context) {
	fmt.Println("calling the Update movie api")
	var movie model.Movie
	movie, error := decodeData(c)
	if error != nil {
		c.JSON(http.StatusInternalServerError, error)
	}

	err := app.ValidateMovie(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	fmt.Println("movie", movie)
	result, err := app.UpdateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) DeleteMovie(c *gin.Context) {
	fmt.Println("calling the Delete movie api")
	id := c.Param("id")
	fmt.Println("movie id", id)
	res, err := app.DeleteMovie(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if res.DeletedCount == 1 {
		c.JSON(http.StatusCreated, "Record Deleted Successfully")
	} else {
		c.JSON(http.StatusAccepted, "Record Not Found")
	}
}

func (*application) GetMovieCount(c *gin.Context) {
	log.Println("calling the get movie count function")
	result, err := app.GetMovieCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, result)
}

func (*application) GetMovieByName(c *gin.Context) {
	movieName := c.Param("moviename")
	result, err := app.GetMovieByName(movieName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func decodeData(c *gin.Context) (model.Movie, error) {
	var movie model.Movie
	fileUrl := cloudinaryConfig(c)
	movie.MovieName = c.Request.FormValue("movieName")
	movie.MovieDetail = c.Request.FormValue("movieDetails")
	movie.MovieTrailer = c.Request.FormValue("trailerLink")
	dateString := c.Request.FormValue("screened_at")
	layout := "2006-01-02"
	date, err := time.Parse(layout, dateString)

	if err != nil {
		return model.Movie{}, err
	}
	date = date.AddDate(0, 0, 1)
	movie.ScreenedAt = date
	movie.MovieImageUrl = fileUrl
	return movie, nil
}

func (*application) GetAllMovies(c *gin.Context) {
	result, err := app.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
