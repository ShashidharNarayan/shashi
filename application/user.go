package application

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (*application) CreateUser(c *gin.Context) {
	fmt.Printf("calling the CreateUser api")
	var user model.User
	c.ShouldBindJSON(&user)

	err := app.ValidateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	result, err := app.CreateUser(user, c) //m.mv.CreateMovie("movie", &movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) Login(c *gin.Context) {
	loggedInObj := c.Request.URL.Query()
	email := loggedInObj.Get("email")
	password := loggedInObj.Get("password")
	result, err := app.Login(email, password, c)
	log.Println("logged In user", result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

/*
checkSession - function to check the user session
*/
func (*application) CheckSession(c *gin.Context) {
	var user model.LoggedInObj
	session := sessions.Default(c)

	email := session.Get("email")
	gender := session.Get("gender")
	is_admin := session.Get("is_admin")
	name := session.Get("name")

	fmt.Println("value details", user)
	fmt.Println("session details", session)
	c.JSON(http.StatusOK, bson.M{"email": email, "gender": gender, "is_admin": is_admin, "name": name})
}

/*
	logout - function to clear user session and logout
*/
func (*application) LogOut(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("email")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete("email")
	session.Delete("gender")
	session.Delete("is_admin")
	session.Delete("name")

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func (*application) GetBookings(c *gin.Context) {
	email := c.Param("email")
	result, err := app.GetBookings(email)
	log.Println("logged In user", result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println("result", result)
	c.JSON(http.StatusAccepted, result)
}

func (*application) UpdateUser(c *gin.Context) {
	log.Println("calling the Update User api")
	var user model.User

	c.ShouldBindJSON(&user)
	log.Println("user", user)

	result, err := app.UpdateUser(user)
	if err != nil {
		log.Println("err", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (*application) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	res, err := app.DeleteUser(id)
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
