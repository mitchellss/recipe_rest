package recipe_rest

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>Test123</h1>
</body>
</html>`

// Holds all the information needed to reconstruct
// a recipe
type Recipe struct {
	Title       string                 `yaml:"title,omitempty"`
	Datetime    int                    `yaml:"datetime,omitempty"`
	Author      string                 `yaml:"author,omitempty"`
	ActiveTime  int                    `yaml:"active_time,omitempty"`
	TotalTime   int                    `yaml:"total_time,omitempty"`
	ServesLow   string                 `yaml:"serves_low,omitempty"`
	ServesHigh  string                 `yaml:"serves_high,omitempty"`
	Ingredients map[string]Measurement `yaml:"ingredients,omitempty"`
	Steps       []Step                 `yaml:"steps,omitempty"`
}

// Holds information about each step in the recipe
type Step struct {
	StepNumber int    `yaml:"step_number,omitempty"`
	StepText   string `yaml:"step_text,omitempty"`
}

// Data type to hold mesurement information.
// i.e. 1.25 cups
type Measurement struct {
	Quantity float32 `yaml:"quantity,omitempty"`
	Unit     string  `yaml:"unit,omitempty"`
}

func NewHandler(r Recipe) http.Handler {
	h := handler{r}
	return h
}

type handler struct {
	r Recipe
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, h.r)
}
