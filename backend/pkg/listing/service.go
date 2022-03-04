package listing

import "errors"

// ErrNotFound is used when a recipe could not be found.
var ErrNotFound = errors.New("recipe not found")

type Service interface {
	GetRecipe(id string) (Recipe, error)
	GetAllRecipes() []Recipe
	GetIngredient(id string) (Ingredient, error)
	GetAllIngredients() []Ingredient
}

type Repository interface {
	GetRecipe(id string) (Recipe, error)
	GetAllRecipes() []Recipe
	GetIngredient(id string) (Ingredient, error)
	GetAllIngredients() []Ingredient
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetRecipe(id string) (Recipe, error) {
	return s.r.GetRecipe(id)
}

func (s *service) GetAllRecipes() []Recipe {
	return s.r.GetAllRecipes()
}

func (s *service) GetIngredient(id string) (Ingredient, error) {
	return s.r.GetIngredient(id)
}

func (s *service) GetAllIngredients() []Ingredient {
	return s.r.GetAllIngredients()
}
