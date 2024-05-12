package pkg

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

func Add(name string) error {
	fmt.Println("Adding task: ", name)

	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = updateDB(db, func(b *bolt.Bucket) error {
		seq, err := b.NextSequence()
		if err != nil {
			return err
		}
		task := Task{
			ID:      TaskID(seq),
			Name:    name,
			Created: time.Now().UTC(),
		}
		encodedTask, err := json.Marshal(&task)
		if err != nil {
			return err
		}
		key := fmt.Sprint(seq)
		return b.Put([]byte(key), encodedTask)
	})
	return err
}
