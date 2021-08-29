package database

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

type task struct {
	number      int
	description string
}

var db *bolt.DB

func SetupDB() {
	var err error
	db, err = bolt.Open("my.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists([]byte("TASKS"))
		if err != nil {
			log.Fatal(err)
		}

		return nil
	})
}

func AddEntry(task string) {
	db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte("TASKS"))
		bucket.Put([]byte("1"), []byte(task))
		return nil
	})
}

func ShowEntries() {
	db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte("TASKS"))
		bucket.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
}

func DeleteEntry(key string) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte("TASKS"))
		err := bucket.Delete([]byte(key))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
