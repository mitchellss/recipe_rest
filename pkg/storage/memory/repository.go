package memory

import (
	"time"

	"github.com/mitchellss/recipe_rest/pkg/service"
)

type Storage struct {
	recipes []Recipe
}

func (m *Storage) AddRecipe(recipe service.Recipe) error {

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

	m.recipes = append(m.recipes, newRecipe)
	return nil
}

func (m *Storage) GetAllRecipes() []service.Recipe {
	var recipes []service.Recipe

	for i := range m.recipes {

		var steps []service.Step
		for i := range m.recipes[i].Steps {
			steps = append(steps, service.Step{
				StepNumber:    m.recipes[i].Steps[i].StepNumber,
				Text:          m.recipes[i].Steps[i].Text,
				IngredientIDs: m.recipes[i].Steps[i].IngredientIDs,
			})
		}

		recipe := service.Recipe{
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

func (m *Storage) GetRecipe(id string) service.Recipe {
	var newRecipe service.Recipe
	for i := range m.recipes {
		if m.recipes[i].ID == id {
			newRecipe.Title = m.recipes[i].Title
			newRecipe.Created = m.recipes[i].Created
			return newRecipe
		}
	}
	return service.Recipe{}
}

// Update method

// Delete method
