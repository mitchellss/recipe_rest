package json

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"path"
	"runtime"
	"time"

	"github.com/mitchellss/recipe_rest/pkg/adding"
	"github.com/mitchellss/recipe_rest/pkg/listing"
	"github.com/mitchellss/recipe_rest/pkg/updating"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/segmentio/ksuid"
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

func (s *Storage) AddRecipe(recipe adding.Recipe) error {

	var steps []Step
	for i := range recipe.Steps {
		steps = append(steps, Step{
			StepNumber:    recipe.Steps[i].StepNumber,
			Text:          recipe.Steps[i].Text,
			IngredientIDs: recipe.Steps[i].IngredientIDs,
		})
	}

	recipeId, err := ksuid.NewRandom()
	if err != nil {
		return err
	}

	newRecipe := Recipe{
		ID:         recipeId.String(),
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func (s *Storage) GetAllRecipes() []listing.Recipe {
	list := []listing.Recipe{}

	records, err := s.db.ReadAll(CollectionRecipe)
	if err != nil {
		return list
	}

	for _, r := range records {
		var jsonRecipe Recipe
		var serviceRecipe listing.Recipe

		if err := json.Unmarshal([]byte(r), &jsonRecipe); err != nil {
			return list
		}

		serviceRecipe.ID = jsonRecipe.ID
		serviceRecipe.Title = jsonRecipe.Title
		serviceRecipe.Author = jsonRecipe.Author
		serviceRecipe.ActiveTime = jsonRecipe.ActiveTime
		serviceRecipe.TotalTime = jsonRecipe.TotalTime
		serviceRecipe.ServesHigh = jsonRecipe.ServesHigh
		serviceRecipe.ServesLow = jsonRecipe.ServesLow
		serviceRecipe.Created = jsonRecipe.Created

		var steps []listing.Step
		for i := range jsonRecipe.Steps {
			steps = append(steps, listing.Step{
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

func (s *Storage) GetRecipe(id string) (listing.Recipe, error) {
	var serviceRecipe listing.Recipe
	var jsonRecipe Recipe

	err := s.db.Read(CollectionRecipe, id, &jsonRecipe)
	if err != nil {
		fmt.Println(err)
		return serviceRecipe, listing.ErrNotFound
	}

	serviceRecipe.ID = jsonRecipe.ID
	serviceRecipe.Title = jsonRecipe.Title
	serviceRecipe.Author = jsonRecipe.Author
	serviceRecipe.ActiveTime = jsonRecipe.ActiveTime
	serviceRecipe.TotalTime = jsonRecipe.TotalTime
	serviceRecipe.ServesHigh = jsonRecipe.ServesHigh
	serviceRecipe.ServesLow = jsonRecipe.ServesLow
	serviceRecipe.Created = jsonRecipe.Created

	var steps []listing.Step
	for i := range jsonRecipe.Steps {
		steps = append(steps, listing.Step{
			StepNumber:    jsonRecipe.Steps[i].StepNumber,
			Text:          jsonRecipe.Steps[i].Text,
			IngredientIDs: jsonRecipe.Steps[i].IngredientIDs,
		})
	}
	serviceRecipe.Steps = steps

	return serviceRecipe, nil
}

func (s *Storage) UpdateRecipe(id string, recipe updating.Recipe) error {
	var jsonRecipe Recipe
	err := s.db.Read(CollectionRecipe, id, &jsonRecipe)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if recipe.Title != "" {
		jsonRecipe.Title = recipe.Title
	}
	if recipe.Author != "" {
		jsonRecipe.Author = recipe.Author
	}
	if recipe.ActiveTime != 0 {
		jsonRecipe.ActiveTime = recipe.ActiveTime
	}
	if recipe.TotalTime != 0 {
		jsonRecipe.TotalTime = recipe.TotalTime
	}
	if recipe.ServesHigh != 0 {
		jsonRecipe.ServesHigh = recipe.ServesHigh
	}
	if recipe.ServesLow != 0 {
		jsonRecipe.ServesLow = recipe.ServesLow
	}
	if recipe.Steps != nil {
		var steps []Step
		for i := range recipe.Steps {
			steps = append(steps, Step{
				StepNumber:    recipe.Steps[i].StepNumber,
				Text:          recipe.Steps[i].Text,
				IngredientIDs: recipe.Steps[i].IngredientIDs,
			})
		}
		jsonRecipe.Steps = steps
	}

	if err := s.db.Write(CollectionRecipe, jsonRecipe.ID, jsonRecipe); err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteRecipe(id string) error {
	return nil
}
