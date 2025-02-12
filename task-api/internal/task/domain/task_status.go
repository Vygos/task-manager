package domain

import (
	"encoding/json"
	"errors"
)

type TaskStatus string

const (
	Completed  TaskStatus = "completed"
	Incomplete TaskStatus = "incomplete"
)

var statuses = []TaskStatus{
	Completed,
	Incomplete,
}

func (s *TaskStatus) String() string {
	return string(*s)
}

func (s *TaskStatus) UnmarshalJSON(b []byte) error {
	var body string
	if err := json.Unmarshal(b, &body); err != nil {
		return err
	}

	*s = TaskStatus(body)

	var find bool

	for _, status := range statuses {
		if status == *s {
			find = true
		}
	}
	if !find {
		return errors.New("invalid task status")
	}

	return nil
}
