package memory

import "time"

type Recipe struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Steps   []Step    `json:"steps"`
}

type Step struct {
	StepNumber int    `json:"step_num"`
	Text       string `json:"text"`
}
