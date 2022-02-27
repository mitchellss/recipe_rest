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

func (s *Storage) AddRecipe(recipe service.Recipe) error {

	var steps []Step
	for i := range recipe.Steps {
		steps = append(steps, Step{
			StepNumber:    recipe.Steps[i].StepNumber,
			Text:          recipe.Steps[i].Text,
			IngredientIDs: recipe.Steps[i].IngredientIDs,
		})
	}

	recipeId := "12345"

	newRecipe := Recipe{
		ID:         recipeId,
		Title:      recipe.Title,
		Author:     recipe.Author,
		ActiveTime: recipe.ActiveTime,
		TotalTime:  recipe.TotalTime,
		ServesHigh: recipe.ServesHigh,
		ServesLow:  recipe.ServesLow,
		Created:    time.Now(),
		Steps:      steps,
	}

	if err := s.db.Write(CollectionRecipe, newRecipe.ID, newRecipe); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetAllRecipes() []service.Recipe {
	list := []service.Recipe{}

	records, err := s.db.ReadAll(CollectionRecipe)
	if err != nil {
		return list
	}

	for _, r := range records {
		var jsonRecipe Recipe
		var serviceRecipe service.Recipe

		if err := json.Unmarshal([]byte(r), &jsonRecipe); err != nil {
			return list
		}

		serviceRecipe.Title = jsonRecipe.Title
		serviceRecipe.Author = jsonRecipe.Author
		serviceRecipe.ActiveTime = jsonRecipe.ActiveTime
		serviceRecipe.TotalTime = jsonRecipe.TotalTime
		serviceRecipe.ServesHigh = jsonRecipe.ServesHigh
		serviceRecipe.ServesLow = jsonRecipe.ServesLow
		serviceRecipe.Created = jsonRecipe.Created

		var steps []service.Step
		for i := range jsonRecipe.Steps {
			steps = append(steps, service.Step{
				StepNumber:    jsonRecipe.Steps[i].StepNumber,
				Text:          jsonRecipe.Steps[i].Text,
				IngredientIDs: jsonRecipe.Steps[i].IngredientIDs,
			})
		}
		serviceRecipe.Steps = steps

		list = append(list, serviceRecipe)
	}
	return list
}

func (s *Storage) GetRecipe(id string) service.Recipe {
	var serviceRecipe service.Recipe

	records, err := s.db.ReadAll(CollectionRecipe)
	if err != nil {
		return serviceRecipe
	}

	for _, r := range records {
		var jsonRecipe Recipe

		if err := json.Unmarshal([]byte(r), &jsonRecipe); err != nil {
			return serviceRecipe
		}

		if jsonRecipe.ID == id {
			serviceRecipe.Title = jsonRecipe.Title
			serviceRecipe.Author = jsonRecipe.Author
			serviceRecipe.ActiveTime = jsonRecipe.ActiveTime
			serviceRecipe.TotalTime = jsonRecipe.TotalTime
			serviceRecipe.ServesHigh = jsonRecipe.ServesHigh
			serviceRecipe.ServesLow = jsonRecipe.ServesLow
			serviceRecipe.Created = jsonRecipe.Created

			var steps []service.Step
			for i := range jsonRecipe.Steps {
				steps = append(steps, service.Step{
					StepNumber:    jsonRecipe.Steps[i].StepNumber,
					Text:          jsonRecipe.Steps[i].Text,
					IngredientIDs: jsonRecipe.Steps[i].IngredientIDs,
				})
			}
			serviceRecipe.Steps = steps

			return serviceRecipe
		}
	}
	return service.Recipe{}
}

// Update method

// Delete method
