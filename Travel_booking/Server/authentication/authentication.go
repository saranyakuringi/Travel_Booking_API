package authentication

import (
	"Travel_Booking/controller"
	"Travel_Booking/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userlogindetails []model.User

func SetAuthenticationConfig() ([]model.User, error) {
	db, err := controller.Connect_DB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		//return nil, err
	}

	users, err := model.User_Query(db)
	if err != nil {
		log.Fatal("Error querying data:", err)
		return nil, err
	}

	userlogindetails = users
	for _, value := range userlogindetails {
		log.Printf("Username: %s, Password: %s", value.Username, value.Password)
	}
	//log.Printf("credentials in userlogindetails- Username:%s, Password:%s",)
	return userlogindetails, err

}

// var hardcodedUsername = "admin"
// var hardcodedPassword = "admin@123"

func AuthMiddleware(userlogindetails []model.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		log.Printf("Received credentials - Username:%s, Password:%s", username, password)
		if ok {
			for _, user := range userlogindetails {
				if username == user.Username && password == user.Password {
					c.Next()
					log.Println("Authentication successful.")
					return
				}
			}
		}
		log.Println("Authentication failed.")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unauthorised"})
		c.Abort()
	}
}
