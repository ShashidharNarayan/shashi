package Service

import (
	"log"
	"regexp"

	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/sha3"
)

var userCollection = "user"
var pattern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (*service) ValidateUser(data *model.User) error {
	if data == nil {
		err := errors.New("User Data Cannot Be Empty")
		return err
	}

	if data.Name == "" {
		err := errors.New("User Name Cannot Be Empty")
		return err
	}

	if data.Email == "" {
		err := errors.New("User Email Cannot Be Empty")
		return err
	}

	if data.Password == "" {
		err := errors.New("User Password Cannot Be Empty")
		return err
	}

	if data.Gender == "" {
		err := errors.New("User Gender Cannot Be Empty")
		return err
	}

	return nil
}

func (*service) ValidateLogin(data *model.User) error {
	if data == nil {
		err := errors.New("User Data Cannot Be Empty")
		return err
	}

	if data.Email == "" {
		err := errors.New("Email Cannot Be Empty")
		return err
	}

	if data.Password == "" {
		err := errors.New("Password Cannot Be Empty")
		return err
	}

	return nil
}

func (*service) Login(email string, password string, c *gin.Context) (model.LoggedInObj, error) {
	if !pattern.MatchString(email) {
		return model.LoggedInObj{}, errors.New("The email field should be a valid email address!")
	}

	buf := []byte(password)
	ciphertext := make([]byte, 64)
	sha3.ShakeSum256(ciphertext, buf)
	encryptptedPassword := string(ciphertext)
	filter := bson.M{"email": email, "password": encryptptedPassword}

	result, err := repo.Login(userCollection, filter)

	if err != nil {
		return model.LoggedInObj{}, err
	}

	session := sessions.Default(c)
	session.Set("email", result.Email)
	session.Set("gender", result.Gender)
	session.Set("is_admin", result.IsAdmin)
	session.Set("name", result.Name)
	session.Save()

	return result, nil
}

func (*service) GetBookings(email string) ([]model.LoggedInObj, error) {
	filter := bson.M{"email": email}
	result, err := repo.GetBookings(userCollection, filter)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) CreateUser(data model.User, c *gin.Context) (interface{}, error) {
	if !pattern.MatchString(data.Email) {
		err := errors.New("The email field should be a valid email address!")
		return nil, err
	}

	data.ID = primitive.NewObjectID()
	filter := bson.M{"email": data.Email}

	buf := []byte(data.Password)
	ciphertext := make([]byte, 64)
	sha3.ShakeSum256(ciphertext, buf)
	data.Password = string(ciphertext)

	result, err := repo.Create(userCollection, data, filter)
	if err != nil {
		return nil, err
	}

	log.Println("create user result", result)

	session := sessions.Default(c)
	session.Set("email", data.Email)
	session.Set("gender", data.Gender)
	session.Set("is_admin", data.IsAdmin)
	session.Set("name", data.Name)
	session.Save()

	return data, nil
}

func (*service) UpdateUser(user model.User) (interface{}, error) {
	filter := bson.M{"email": user.Email}

	var data bson.M
	if user.IsAdmin {
		data = bson.M{
			"$set": bson.M{
				"email":       user.Email,
				"is_admin":    user.IsAdmin,
				"my_bookings": user.MyBookings,
			},
		}
	} else {
		data = bson.M{
			"$set": bson.M{
				"email":       user.Email,
				"my_bookings": user.MyBookings,
			},
		}
	}

	result, err := repo.Update(userCollection, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) DeleteUser(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	result, err := repo.Delete(userCollection, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// func (*service) GetUser() ([]model.User, error)
