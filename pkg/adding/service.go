package adding

type Service interface {
	AddRecipe(...Recipe) error
	// AddIngredient(...Ingredient) error
}

type Repository interface {
	AddRecipe(Recipe) error
	// AddIngredient(Ingredient) error
}

func NewService(r Repository) Service {
	return &service{r}
}

type service struct {
	r Repository
}

func (s *service) AddRecipe(recipe ...Recipe) error {
	for _, rr := range recipe {
		_ = s.r.AddRecipe(rr)
	}
	return nil
}
