package database

import (
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

func TestShowTasks(t *testing.T) {
	ShowTasks()
}

func TestCompleteTask(t *testing.T) {
	CompleteTask("test")
}
