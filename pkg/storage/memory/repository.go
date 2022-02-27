package memory

import (
	"time"

	"github.com/mitchellss/recipe_rest/pkg/adding"
	"github.com/mitchellss/recipe_rest/pkg/listing"
	"github.com/mitchellss/recipe_rest/pkg/storage"
)

type Storage struct {
	recipes []Recipe
}

func (m *Storage) AddRecipe(recipe adding.Recipe) error {

	var steps []Step
	for i := range recipe.Steps {
		steps = append(steps, Step{
			StepNumber:    recipe.Steps[i].StepNumber,
			Text:          recipe.Steps[i].Text,
			IngredientIDs: recipe.Steps[i].IngredientIDs,
		})
	}

	recipeId, err := storage.GenerateID()
	if err != nil {
		return err
	}

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

	m.recipes = append(m.recipes, newRecipe)
	return nil
}

func (m *Storage) GetAllRecipes() []listing.Recipe {
	var recipes []listing.Recipe

	for i := range m.recipes {

		var steps []listing.Step
		for i := range m.recipes[i].Steps {
			steps = append(steps, listing.Step{
				StepNumber:    m.recipes[i].Steps[i].StepNumber,
				Text:          m.recipes[i].Steps[i].Text,
				IngredientIDs: m.recipes[i].Steps[i].IngredientIDs,
			})
		}

		recipe := listing.Recipe{
			ID:         m.recipes[i].ID,
			Title:      m.recipes[i].Title,
			Author:     m.recipes[i].Author,
			ActiveTime: m.recipes[i].ActiveTime,
			TotalTime:  m.recipes[i].TotalTime,
			ServesHigh: m.recipes[i].ServesHigh,
			ServesLow:  m.recipes[i].ServesLow,
			Created:    m.recipes[i].Created,
			Steps:      steps,
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

func (m *Storage) GetRecipe(id string) (listing.Recipe, error) {
	for i := range m.recipes {
		if m.recipes[i].ID == id {
			var steps []listing.Step
			for i := range m.recipes[i].Steps {
				steps = append(steps, listing.Step{
					StepNumber:    m.recipes[i].Steps[i].StepNumber,
					Text:          m.recipes[i].Steps[i].Text,
					IngredientIDs: m.recipes[i].Steps[i].IngredientIDs,
				})
			}
			newRecipe := listing.Recipe{
				ID:         m.recipes[i].ID,
				Title:      m.recipes[i].Title,
				Author:     m.recipes[i].Author,
				ActiveTime: m.recipes[i].ActiveTime,
				TotalTime:  m.recipes[i].TotalTime,
				ServesHigh: m.recipes[i].ServesHigh,
				ServesLow:  m.recipes[i].ServesLow,
				Created:    m.recipes[i].Created,
				Steps:      steps,
			}
			return newRecipe, nil
		}
	}
	return listing.Recipe{}, listing.ErrNotFound
}

// func (m *Storage) UpdateRecipe(id string, service.Recipe) {}

// func (m *Storage) DeleteRecipe(id string) {}
