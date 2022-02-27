package listing

import "time"

type Recipe struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	ActiveTime int       `json:"active_time"`
	TotalTime  int       `json:"total_time"`
	ServesHigh int       `json:"serves_high"`
	ServesLow  int       `json:"serves_low"`
	Created    time.Time `json:"created"`
	Steps      []Step    `json:"steps"`
}

type Step struct {
	StepNumber    int      `json:"step_num"`
	Text          string   `json:"text"`
	IngredientIDs []string `json:"ingredient_ids"`
}
