package service

type Service interface {
	AddRecipe(...Recipe) error
	GetRecipe(id string) Recipe
	GetAllRecipes() []Recipe
}

type Repository interface {
	AddRecipe(Recipe) error
	GetRecipe(id string) Recipe
	GetAllRecipes() []Recipe
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddRecipe(recipe ...Recipe) error {
	for _, rr := range recipe {
		_ = s.r.AddRecipe(rr)
	}
	return nil
}

func (s *service) GetRecipe(id string) Recipe {
	return s.r.GetRecipe(id)
}

func (s *service) GetAllRecipes() []Recipe {
	return s.r.GetAllRecipes()
}
