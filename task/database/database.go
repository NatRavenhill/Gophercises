package database

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

var Database *bolt.DB

func SetupDB() {
	Database, err := bolt.Open("my.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer Database.Close()
}
