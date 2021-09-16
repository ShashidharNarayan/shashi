package Repository

import "github.com/ShashidharNarayan/bms/repository/model"

func (*repo) GetShowsOfTheMovie(collection string, movieName string) ([]model.MovieShow, error) {
	result, err := GetShowsOfTheMovie(collection, movieName)

	if err != nil {
		return nil, err
	}

	return result, nil
}
