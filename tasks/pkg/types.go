package pkg

import "time"

type TaskID uint

type Task struct {
	id        TaskID
	name      string
	created   time.Time
	completed time.Time
}
