package api

import (
	"Travel_Booking/controller"
	"Travel_Booking/model"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_CustomerBookings(c *gin.Context) {
	id, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(404, gin.H{"message": "Error in request body"})
		return
	}

	var data model.Booking
	err = json.Unmarshal(id, &data)
	if err != nil {
		log.Println("Error in unmarshal data", err)
		c.IndentedJSON(404, gin.H{"message": "Error in unmarshal data"})
		return
	}

	db, err := controller.Connect_DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	bookings, err := model.Query(db, data.Custid, data.Lastname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute the query"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func GET_NewBooking(c *gin.Context) {
	id, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(404, gin.H{"message": "Error in request body"})
		return
	}
	var Newdata model.Customer
	err = json.Unmarshal(id, &Newdata)
	if err != nil {
		log.Println("Error in unmarshal data", err)
		c.IndentedJSON(404, gin.H{"message": "Error in unmarshal data"})
		return
	}

	db, err := controller.Connect_DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	NewBooking, err := model.User_Insert(db, Newdata.Custid, Newdata.Firstname, Newdata.Lastname, Newdata.Address, Newdata.City, Newdata.Tourid, Newdata.Payment_amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute the query"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"bookings": NewBooking})
	}
}
