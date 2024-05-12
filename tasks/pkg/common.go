package pkg

import (
	"os"
	"path"
	"time"

	"github.com/boltdb/bolt"
)

type TaskID uint64

type Task struct {
	ID        TaskID    `json:"id"`
	Name      string    `json:"name"`
	Created   time.Time `json:"created"`
	Completed time.Time `json:"completed"`
}

func openDB() (*bolt.DB, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath := path.Join(homeDir, "tasks.db")
	db, err := bolt.Open(filePath, 0600, nil)
	return db, err
}

const bucketName = "Tasks"

func updateDB(db *bolt.DB, callback func(b *bolt.Bucket) error) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return callback(b)
	})
}

func viewDB(db *bolt.DB, callback func(b *bolt.Bucket) error) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return callback(b)
	})
}
