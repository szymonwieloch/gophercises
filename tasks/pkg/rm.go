package pkg

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func Rm(id TaskID) error {
	fmt.Println("Removing task", id)
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()
	return updateDB(db, func(b *bolt.Bucket) error {
		key := fmt.Sprint(id)
		task := b.Get([]byte(key))
		if task == nil {
			return fmt.Errorf("Task %d was not found", id)
		}
		return b.Delete([]byte(key))
	})
}
