package database

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

func GetTasksBucket(t *bolt.Tx) *bolt.Bucket {
	bucket := t.Bucket([]byte("TASKS"))
	if bucket == nil {
		log.Fatal("Tasks bucket not found")
	}
	return bucket
}

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

func AddTask(task string) {
	db.Update(func(t *bolt.Tx) error {
		bucket := GetTasksBucket(t)
		total := bucket.Stats().KeyN
		bucket.Put([]byte(fmt.Sprintf("%d", total+1)), []byte(task))

		fmt.Printf("Added \"%s\" to your task list.", task)
		return nil
	})
}

func ShowTasks() {
	db.View(func(t *bolt.Tx) error {
		bucket := GetTasksBucket(t)
		bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("%s. %s \n", string(k), string(v))
			return nil
		})
		return nil
	})
}

func CompleteTask(key string) {
	var value []byte
	err := db.Update(func(t *bolt.Tx) error {
		bucket := GetTasksBucket(t)
		value = bucket.Get([]byte(key))

		err := bucket.Delete([]byte(key))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You have completed the \"%s\" task.\n", value)

}
