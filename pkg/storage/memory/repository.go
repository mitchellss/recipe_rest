package memory

import (
	"time"

	"github.com/mitchellss/recipe_rest/pkg/service"
)

type Storage struct {
	recipes []Recipe
}

func (m *Storage) AddRecipe(recipe service.Recipe) error {
	newRecipe := Recipe{
		ID:      recipe.ID,
		Name:    recipe.Name,
		Created: time.Now(),
	}
	m.recipes = append(m.recipes, newRecipe)
	return nil
}

func (m *Storage) GetRecipe(id string) service.Recipe {
	var newRecipe service.Recipe
	for i := range m.recipes {
		if m.recipes[i].ID == id {
			newRecipe.ID = m.recipes[i].ID
			newRecipe.Name = m.recipes[i].Name
			newRecipe.Created = m.recipes[i].Created
			return newRecipe
		}
	}
	return service.Recipe{}
}

func (m *Storage) GetAllRecipes() []service.Recipe {
	var recipes []service.Recipe
	for i := range m.recipes {
		recipe := service.Recipe{
			ID:      m.recipes[i].ID,
			Name:    m.recipes[i].Name,
			Created: m.recipes[i].Created,
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}
