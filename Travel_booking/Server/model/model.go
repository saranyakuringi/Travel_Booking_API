package model

import (
	"database/sql"
	"log"
)

type Customer struct {
	Custid         int     `json:"custid"`
	Firstname      string  `json:"firstname"`
	Lastname       string  `json:"lastname"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	Tourid         string  `json:"tourid"`
	Payment_amount float32 `json:"payment_amount"`
}

type Tourid struct {
	Tourid   string `json:"tourid"`
	Tourname string `json:"tourname"`
	Tourdate string `json:"tourdata"`
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

func Query(db *sql.DB, custid int, lastname string) ([]Booking, error) {
	rows, err := db.Query("Select * From Booking where custid=$1 OR lastname=$2", custid, lastname)
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
	return output, err
}

func User_Insert(db *sql.DB, custid int, firstname, lastname, address, city, tourid string, payment_amount float32) ([]Booking, error) {
	// Inserting values in the customer table
	_, err := db.Exec("INSERT INTO customer (custid, firstname, lastname, address, city, tourid, payment_amount) VALUES ($1, $2, $3, $4, $5, $6, $7)", custid, firstname, lastname, address, city, tourid, payment_amount)
	if err != nil {
		log.Println("Error inserting into customer table:", err)
		return nil, err
	}

	// Select values from tour table
	rows_tour, err := db.Query("SELECT * FROM tour WHERE tourid = $1", tourid)
	if err != nil {
		log.Println("Error in tour select query:", err)
		return nil, err
	}
	defer rows_tour.Close()

	var tour []Tourid
	for rows_tour.Next() {
		var tourdata Tourid
		err := rows_tour.Scan(&tourdata.Tourid, &tourdata.Tourname, &tourdata.Tourdate)
		if err != nil {
			log.Println("Error in tour scan:", err)
			return nil, err
		}
		tour = append(tour, tourdata)
	}

	if err := rows_tour.Err(); err != nil {
		log.Println("Error in tour rows iteration:", err)
		return nil, err
	}

	// Inserting values into the booking table
	_, err = db.Exec("INSERT INTO booking (custid, firstname, lastname, tourid, tourname, tourdate, payment_amount) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		custid, firstname, lastname, tourid, tour[0].Tourname, tour[0].Tourdate, payment_amount)
	if err != nil {
		log.Println("Error inserting into booking table:", err)
		return nil, err
	}

	// Fetch the inserted booking record
	rows_booking, err := db.Query("SELECT * FROM booking WHERE custid = $1", custid)
	if err != nil {
		log.Println("Error fetching booking data:", err)
		return nil, err
	}
	defer rows_booking.Close()

	var booking []Booking
	for rows_booking.Next() {
		var bookingdata Booking
		err := rows_booking.Scan(&bookingdata.Custid, &bookingdata.Firstname, &bookingdata.Lastname, &bookingdata.Tourid, &bookingdata.Tourname, &bookingdata.Tourdate, &bookingdata.Payment_amount)
		if err != nil {
			log.Println("Error in booking scan:", err)
			return nil, err
		}
		booking = append(booking, bookingdata)
	}

	if err := rows_booking.Err(); err != nil {
		log.Println("Error in booking rows iteration:", err)
		return nil, err
	}

	return booking, nil
}
