package model

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type Customer struct {
	Custid         int     `json:"custid"`
	Firstname      string  `json:"firstname"`
	Lastname       string  `json:"lastname"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	Tourid         int     `json:"tourid"`
	Payment_amount float32 `json:"payment_amount"`
}

type Tourid struct {
	Tourid   string      `json:"tourid"`
	Tourname string      `json:"tourname"`
	Tourdate pq.NullTime `json:"tourdata"`
}

type Booking struct {
	Custid         int     `json:"custid"`
	Firstname      string  `json:"firstname"`
	Lastname       string  `json:"lastname"`
	Tourid         string  `json:"tourid"`
	Tourname       string  `json:"tourname"`
	Tourdate       string  `json:"tourdata"`
	Payment_amount float32 `json:"payment_amount"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Query(db *sql.DB, custid int) ([]Booking, error) {
	rows, err := db.Query("Select * From Booking where custid=$1", custid)
	if err != nil {
		log.Println("Error in rows query", err)
	}
	var bookingData Booking
	var output []Booking

	for rows.Next() {
		err = rows.Scan(
			&bookingData.Custid,
			&bookingData.Firstname,
			&bookingData.Lastname,
			&bookingData.Tourid,
			&bookingData.Tourname,
			&bookingData.Tourdate,
			&bookingData.Payment_amount,
		)
		if err != nil {
			log.Println("Error in err", err)
		}

		output = append(output, bookingData)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in rows iteration", err)
		return nil, err
	}

	return output, nil
}

func User_Query(db *sql.DB) ([]User, error) {
	rows, err := db.Query("Select * From \"user\"")
	if err != nil {
		log.Println("Error in rows query", err)
	}
	defer rows.Close()

	var output []User

	for rows.Next() {
		var userdata User
		err := rows.Scan(&userdata.Username, &userdata.Password)
		if err != nil {
			log.Println("Error in err", err)
		}
		output = append(output, userdata)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error in rows iteration", err)
		return nil, err
	}
	return output, nil

}
