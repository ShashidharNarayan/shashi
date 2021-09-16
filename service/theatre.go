package Service

import (
	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var theatreCollection = "theatre"

func (*service) ValidateTheatre(data *model.Theatre) error {
	if data == nil {
		err := errors.New("Theatre Data Cannot Be Empty")
		return err
	}

	if data.TheatreName == "" {
		err := errors.New("Theatre Name Cannot Be Empty")
		return err
	}

	if data.TheatreUrl == "" {
		err := errors.New("Theatre Start Image Url Cannot Be Empty")
		return err
	}

	return nil
}

func (*service) GetTheatres() ([]model.Theatre, error) {
	filter := bson.M{}
	result, err := repo.GetTheatres(theatreCollection, filter)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) CreateTheatre(data model.Theatre) (interface{}, error) {
	data.ID = primitive.NewObjectID()
	filter := bson.M{"theatre_name": data.TheatreName}
	result, err := repo.Create(theatreCollection, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) UpdateTheatre(theatre model.Theatre) (interface{}, error) {
	filter := bson.M{"_id": theatre.ID}
	data := bson.M{
		"$set": bson.M{
			"theatre_name": theatre.TheatreName,
			"theatre_url":  theatre.TheatreUrl,
		},
	}

	result, err := repo.Update(theatreCollection, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) DeleteTheatre(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	result, err := repo.Delete(theatreCollection, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
