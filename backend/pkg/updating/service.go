package updating

import (
	"errors"
)

// ErrNotFound is used when a recipe could not be found.
var ErrNotFound = errors.New("recipe not found")

type Service interface {
	UpdateRecipe(id string, recipe Recipe) error
	UpdateIngredient(id string, ingredient Ingredient) error
}

type Repository interface {
	UpdateRecipe(id string, recipe Recipe) error
	UpdateIngredient(id string, ingredient Ingredient) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateRecipe(id string, recipe Recipe) error {
	return s.r.UpdateRecipe(id, recipe)
}

func (s *service) UpdateIngredient(id string, ingredient Ingredient) error {
	return s.r.UpdateIngredient(id, ingredient)
}
