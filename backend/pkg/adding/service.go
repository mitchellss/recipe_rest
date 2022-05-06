package adding

type Service interface {
	AddRecipe(...Recipe) error
	AddIngredient(...Ingredient) error
	AddUnit(...UnitDict) error
}

type Repository interface {
	AddRecipe(Recipe) error
	AddIngredient(Ingredient) error
	AddUnit(unit string, units_per_cup float64) error
}

func NewService(r Repository) Service {
	return &service{r}
}

type service struct {
	r Repository
}

func (s *service) AddRecipe(recipe ...Recipe) error {
	// TODO: Add check for duplicate
	// TODO: Add check for ingredient reference for ingredient_ids field
	for _, rr := range recipe {
		_ = s.r.AddRecipe(rr)
	}
	return nil
}

func (s *service) AddIngredient(ingredient ...Ingredient) error {
	// TODO: Add check for duplicate
	// TODO: Add check for ingredient reference in substitutes field
	for _, rr := range ingredient {
		_ = s.r.AddIngredient(rr)
	}
	return nil
}

func (s *service) AddUnit(unit ...UnitDict) error {
	for _, rr := range unit {
		for key, val := range rr.Dict {
			_ = s.r.AddUnit(key, val)
		}
	}
	return nil
}
