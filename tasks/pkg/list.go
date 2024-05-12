package pkg

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gosuri/uitable"
)

func List(open, completed bool) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()
	return viewDB(db, func(b *bolt.Bucket) error {
		if b == nil {
			return nil
		}
		c := b.Cursor()
		table := uitable.New()
		table.AddRow("ID", "NAME", "CREATED", "COMPLETED")
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task Task
			err = json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			completed := "---"
			var nullTime time.Time
			if task.Completed != nullTime {
				completed = task.Completed.Local().Format(time.DateTime)
			}
			table.AddRow(task.ID, task.Name, task.Created.Local().Format(time.DateTime), completed)
		}
		fmt.Println(table)
		return nil
	})
}
