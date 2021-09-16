package Repository

import (
	"github.com/ShashidharNarayan/bms/repository/model"
)

func (*repo) GetMovieCoupons(collection string) ([]model.Coupon, error) {
	result, err := GetCoupons(collection)

	if err != nil {
		return nil, err
	}

	return result, nil
}
