package application

import (
	"net/http"

	"github.com/ShashidharNarayan/bms/repository/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const couponCollection = "coupons"

func (*application) CreateCoupons(c *gin.Context) {
	var coupon model.Coupon
	c.ShouldBindJSON(&coupon)
	err := app.ValidateCoupon(&coupon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	coupon.ID = primitive.NewObjectID()
	result, err := app.CreateCoupon(coupon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) GetCoupons(c *gin.Context) {
	result, err := app.GetCoupon()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) UpdateCoupon(c *gin.Context) {
	var coupon model.Coupon
	c.ShouldBindJSON(&coupon)

	err := app.ValidateCoupon(&coupon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	result, err := app.UpdateCoupon(coupon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (*application) DeleteCoupons(c *gin.Context) {
	id := c.Param("id")
	res, err := app.DeleteCoupon(id)
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
