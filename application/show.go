package application

import (
	"log"
	"net/http"

	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/gin-gonic/gin"
)

func (*application) GetShows(c *gin.Context) {
	movieName := c.Param("moviename")
	result, err := app.GetShowsOfTheMovie(movieName) //m.mv.CreateMovie("movie", &movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) CreateShow(c *gin.Context) {
	log.Println("calling the create show api")
	var show model.Show
	c.ShouldBindJSON(&show)

	err := app.ValidateShow(&show)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	result, err := app.CreateShow(show) //m.mv.CreateMovie("movie", &movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) UpdateShow(c *gin.Context) {
	var show model.Show
	c.ShouldBindJSON(&show)

	err := app.ValidateShow(&show)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	result, err := app.UpdateShow(show)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) DeleteShow(c *gin.Context) {
	id := c.Param("id")
	res, err := app.DeleteShow(id)
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
