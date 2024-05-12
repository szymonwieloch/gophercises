package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gosuri/uitable"
)

// Adds a task to the database
func Add(name string) error {
	fmt.Println("Adding task:", name)

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
		key := dbID(task.ID)
		return b.Put([]byte(key), encodedTask)
	})
	return err
}

// Lists tasks from the database
func List(showOpen, showCompleted bool) error {
	taskKinds := []string{}
	if showOpen {
		taskKinds = append(taskKinds, "open")
	}
	if showCompleted {
		taskKinds = append(taskKinds, "completed")
	}
	fmt.Println("Listing tasks:", strings.Join(taskKinds, ", "))
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()
	var nullTime time.Time
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
			isCompleted := task.Completed != nullTime
			if !showCompleted && isCompleted {
				continue
			}
			if !showOpen && !isCompleted {
				continue
			}
			completedText := "---"
			if isCompleted {
				completedText = task.Completed.Local().Format(time.DateTime)
			}
			table.AddRow(task.ID, task.Name, task.Created.Local().Format(time.DateTime), completedText)
		}
		fmt.Println(table)
		return nil
	})
}

// Removes a task from the database
func Rm(id TaskID) error {
	fmt.Println("Removing task", id)
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()
	return updateDB(db, func(b *bolt.Bucket) error {
		key := dbID(id)
		task := b.Get([]byte(key))
		if task == nil {
			return fmt.Errorf("Task %d was not found", id)
		}
		return b.Delete([]byte(key))
	})
}

// Completes the task in the database
func Complete(id TaskID) error {
	fmt.Println("Completing task", id)
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()
	return updateDB(db, func(b *bolt.Bucket) error {
		key := dbID(id)
		data := b.Get(key)
		if data == nil {
			return fmt.Errorf("Task %d not found", id)
		}
		var task Task
		err = json.Unmarshal(data, &task)
		if err != nil {
			return err
		}
		var nullTime time.Time
		if task.Completed != nullTime {
			return fmt.Errorf("Task %d is already completed", id)
		}
		task.Completed = time.Now().UTC()
		data, err = json.Marshal(&task)
		if err != nil {
			return err
		}
		return b.Put(key, data)
	})
}
