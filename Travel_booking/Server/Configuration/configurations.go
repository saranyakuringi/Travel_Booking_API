package Configuration

import (
	"encoding/json"
	"log"
	"os"
)

// database structure
type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

// Authentication structure
type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Read data from config.json file to connect to db
func LoadDatabase(filename string) (*Database, error) {

	//file open
	filepath, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Error in opening the config file", err)
		return nil, err
	}

	// read the contents
	var database_connections Database
	err = json.Unmarshal(filepath, &database_connections)
	if err != nil {
		log.Println("Error in opening the unmarshal", err)
		return nil, err
	}
	return &database_connections, err
}
