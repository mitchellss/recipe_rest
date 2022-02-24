package json

import (
	"encoding/json"
	"path"
	"runtime"
	"time"

	"github.com/mitchellss/recipe_rest/pkg/service"
	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	dir              = "/data/"
	CollectionRecipe = "recipes"
)

// Storage stores beer data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) GetAllRecipes() []service.Recipe {
	list := []service.Recipe{}

	records, err := s.db.ReadAll(CollectionRecipe)
	if err != nil {
		return list
	}

	for _, r := range records {
		var recipe1 Recipe
		var recipe2 service.Recipe

		if err := json.Unmarshal([]byte(r), &recipe1); err != nil {
			return list
		}

		recipe2.ID = recipe1.ID
		recipe2.Created = recipe1.Created
		recipe2.Name = recipe1.Name

		var steps []service.Step
		for i := range recipe1.Steps {
			steps = append(steps, service.Step{
				StepNumber: recipe1.Steps[i].StepNumber,
				Text:       recipe1.Steps[i].Text,
			})
		}
		recipe2.Steps = steps

		list = append(list, recipe2)
	}
	return list
}

func (s *Storage) AddRecipe(recipe service.Recipe) error {

	var steps []Step
	for i := range recipe.Steps {
		steps = append(steps, Step{
			StepNumber: recipe.Steps[i].StepNumber,
			Text:       recipe.Steps[i].Text,
		})
	}

	newRecipe := Recipe{
		ID:      recipe.ID,
		Name:    recipe.Name,
		Created: time.Now(),
		Steps:   steps,
	}

	if err := s.db.Write(CollectionRecipe, newRecipe.ID, newRecipe); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetRecipe(id string) service.Recipe {
	return service.Recipe{}
}
