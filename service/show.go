package Service

import (
	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName string = "show"

func (*service) ValidateShow(data *model.Show) error {
	if data == nil {
		err := errors.New("Show Data Cannot Be Empty")
		return err
	}

	if data.MovieName == "" {
		err := errors.New("Show Name Cannot Be Empty")
		return err
	}

	if data.MovieStartTime == "" {
		err := errors.New("Show Start Time Cannot Be Empty")
		return err
	}

	if data.ShowDate.IsZero() {
		err := errors.New("Show Trailer Cannot Be Empty")
		return err
	}

	return nil
}

func (*service) GetShowsOfTheMovie(movieName string) ([]model.MovieShow, error) {
	result, err := repo.GetShowsOfTheMovie(collectionName, movieName)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) CreateShow(data model.Show) (interface{}, error) {
	data.ID = primitive.NewObjectID()
	filter := bson.M{"movie_name": data.MovieName, "theatre_name": data.TheatreName, "screened_at": data.ShowDate, "show_time": data.StartTime}
	result, err := repo.Create(collectionName, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) UpdateShow(show model.Show) (interface{}, error) {

	filter := bson.M{"_id": show.ID}
	data := bson.M{
		"$set": bson.M{
			"seats": show.Seats,
		},
	}

	result, err := repo.Update(collectionName, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) DeleteShow(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	result, err := repo.Delete(collectionName, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
