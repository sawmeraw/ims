package models

type TaskFrequency int

const (
	Daily TaskFrequency = iota
	Weekly
	Fortnightly
	Monthly
)

func (t TaskFrequency) String() string {
	return [...]string{"Daily", "Weekly", "Fortnightly", "Monthly"}[t]
}

type Task struct {
	ID             int
	ClientID       int
	UserID         int
	Description    string
	TaskFrequency  TaskFrequency
	FrequencyCount int
}
