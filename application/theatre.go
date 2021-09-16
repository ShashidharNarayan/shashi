package application

import (
	"log"
	"net/http"

	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/gin-gonic/gin"
)

func (*application) CreateTheatre(c *gin.Context) {
	log.Println("calling the create theatre api")
	var theatre model.Theatre
	fileUrl := cloudinaryConfig(c)
	theatre.TheatreName = c.Request.FormValue("theatreName")
	theatre.TheatreUrl = fileUrl

	err := app.ValidateTheatre(&theatre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	result, err := app.CreateTheatre(theatre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) GetTheatres(c *gin.Context) {
	log.Println("calling the getMovie api")
	result, err := app.GetTheatres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) UpdateTheatre(c *gin.Context) {
	var theatre model.Theatre
	c.ShouldBindJSON(&theatre)

	err := app.ValidateTheatre(&theatre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	result, err := app.UpdateTheatre(theatre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) DeleteTheatre(c *gin.Context) {
	id := c.Param("id")
	res, err := app.DeleteTheatre(id)
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
