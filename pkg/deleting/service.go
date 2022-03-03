package deleting

import "errors"

// ErrNotFound is used when a recipe could not be found.
var ErrNotFound = errors.New("recipe not found")

type Service interface {
	DeleteRecipe(id string) error
	DeleteIngredient(id string) error
}

type Repository interface {
	DeleteRecipe(id string) error
	DeleteIngredient(id string) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteRecipe(id string) error {
	return s.r.DeleteRecipe(id)
}

func (s *service) DeleteIngredient(id string) error {
	return s.r.DeleteIngredient(id)
}
