package database

import (
	"bytes"
	"io/ioutil"
	"testing"

	bolt "go.etcd.io/bbolt"
)

func init() {
	//check bucket exists if so delete all stuff else create db
	f, _ := ioutil.TempFile("", "")
	path := f.Name()
	f.Close()
	SetupDB(path)
}

func TestGetTasksBucket(t *testing.T) {
	db.View(func(tx *bolt.Tx) error {
		GetTasksBucket(tx)
		return nil
	})
}

func TestAddTask(t *testing.T) {
	db.View(func(tx *bolt.Tx) error {
		AddTask("test task")
		bucket := GetTasksBucket(tx)
		actualValue := bucket.Get([]byte("1"))

		if !bytes.Equal(actualValue, []byte("test task")) {
			t.Fatalf("Got %s, expected test task", actualValue)
		}

		return nil
	})

}
