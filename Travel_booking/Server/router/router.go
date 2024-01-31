package router

import (
	"Travel_Booking/api"
	"Travel_Booking/authentication"
	"Travel_Booking/model"

	"github.com/gin-gonic/gin"
)

func Router(users []model.User) {
	router := gin.Default()
	//router.GET("/Travelbooking", api.Get_Query)
	router.GET("/Travelbooking/bookings/Search/custid/", authentication.AuthMiddleware(users), func(c *gin.Context) {
		api.Get_CustomerBookings(c)
	})
	router.GET("/Travelbooking/bookings/NewBooking", authentication.AuthMiddleware(users), func(c *gin.Context) {
		api.GET_NewBooking(c)
	})
	router.Run("localhost:8090")

}

//
