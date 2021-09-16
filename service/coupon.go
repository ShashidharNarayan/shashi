package Service

import (
	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var couponCollection = "coupon"

func (*service) CreateCoupon(data model.Coupon) (interface{}, error) {
	data.ID = primitive.NewObjectID()
	filter := bson.M{"CouponCode": data.CouponCode}
	result, err := repo.Create(couponCollection, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) UpdateCoupon(coupon model.Coupon) (interface{}, error) {
	filter := bson.M{"_id": coupon.ID}
	data := bson.M{
		"$set": bson.M{
			"CouponCode": coupon.CouponCode,
		},
	}

	result, err := repo.Update(couponCollection, data, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (*service) DeleteCoupon(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	result, err := repo.Delete(couponCollection, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) GetCoupon() ([]model.Coupon, error) {
	result, err := repo.GetMovieCoupons(couponCollection)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*service) ValidateCoupon(data *model.Coupon) error {
	if data == nil {
		err := errors.New("Coupon Data Is Empty")
		return err
	}

	if data.CouponCode == "" {
		err := errors.New("Coupon Code Cannot Be Empty")
		return err
	}

	if data.CouponPrice == 0 {
		err := errors.New("Coupon Price Cannot Be Zero")
		return err
	}

	return nil
}
