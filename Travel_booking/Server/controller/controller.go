package controller

import (
	"Travel_Booking/Configuration"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func Connect_DB() (*sql.DB, error) {

	filepath, err := Configuration.LoadDatabase("config.json")
	if err != nil {
		log.Println("Error in opening the file", err)
	}
	dbconfig := Configuration.Database{
		Host:     filepath.Host,
		Port:     filepath.Port,
		User:     filepath.User,
		Password: filepath.Password,
		Dbname:   filepath.Dbname,
	}

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable", dbconfig.Host, dbconfig.Port, dbconfig.User, dbconfig.Password, dbconfig.Dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		log.Println("Error in connecting to database", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}
	log.Println("Connected to database sucessfully!!!")

	return db, err

}
