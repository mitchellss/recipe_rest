package adding

type Recipe struct {
	Title      string     `json:"title"`
	Author     string     `json:"author"`
	ActiveTime int        `json:"active_time"`
	TotalTime  int        `json:"total_time"`
	ServesHigh int        `json:"serves_high"`
	ServesLow  int        `json:"serves_low"`
	Steps      []Step     `json:"steps"`
	Materials  []Material `json:"materials"`
}

type Step struct {
	StepNumber    int      `json:"step_num"`
	Text          string   `json:"text"`
	IngredientIDs []string `json:"ingredient_ids"`
}

type Material struct {
	MaterialNumber int     `json:"material_num"`
	IngredientID   string  `json:"ingredient_id"`
	Unit           string  `json:"unit"`
	Amount         float64 `json:"amount"`
	Quality        string  `json:"quality"`
	Note           string  `json:"note"`
	Optional       bool    `json:"optional"`
}
