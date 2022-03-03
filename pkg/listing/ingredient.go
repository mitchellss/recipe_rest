package listing

type Ingredient struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	MassInGrams int    `json:"mass_in_grams"`
	Substitutes string `json:"substitutes"`
}
