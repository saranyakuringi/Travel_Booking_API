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

// func Get_Query(c *gin.Context) {
// 	input, err := model.Output()
// 	if err != nil {
// 		log.Println("Error in model output:", err)
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Message:": "Error in model output"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, gin.H{"Output": input})

// }

func Get_CustomerBookings(c *gin.Context) {
	//id := c.Param("tourid")
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

	bookings, err := model.Query(db, data.Custid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute the query"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}
