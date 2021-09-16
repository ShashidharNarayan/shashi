package Service

import (
	"testing"

	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/stretchr/testify/assert"
)

func TestValiddation(t *testing.T) {
	testService := NewService()
	err := testService.ValidateCoupon(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Coupon Data Is Empty", err.Error())
}

func TestValiddationTitle(t *testing.T) {
	post := model.Coupon{CouponCode: "", CouponPrice: 100}

	testService := NewService()

	err := testService.ValidateCoupon(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "Coupon Code Cannot Be Epty", err.Error())
}
