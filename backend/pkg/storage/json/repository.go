package json

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/mitchellss/recipe_rest/pkg/adding"
	"github.com/mitchellss/recipe_rest/pkg/deleting"
	"github.com/mitchellss/recipe_rest/pkg/listing"
	"github.com/mitchellss/recipe_rest/pkg/updating"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/segmentio/ksuid"
)

const (
	dir                  = "../../../../data/"
	CollectionRecipe     = "recipes"
	CollectionIngredient = "ingredients"
	CollectionUnits      = "units"
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

	var materials []Material
	for i := range recipe.Materials {
		materials = append(materials, Material{
			MaterialNumber: recipe.Materials[i].MaterialNumber,
			IngredientID:   recipe.Materials[i].IngredientID,
			Unit:           recipe.Materials[i].Unit,
			Amount:         recipe.Materials[i].Amount,
			Quality:        recipe.Materials[i].Quality,
			Note:           recipe.Materials[i].Note,
			Optional:       recipe.Materials[i].Optional,
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
		Materials:  materials,
	}

	if err := s.db.Write(CollectionRecipe, newRecipe.ID, newRecipe); err != nil {
		return err
	}
	return nil
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

		var materials []listing.Material
		for i := range jsonRecipe.Materials {
			materials = append(materials, listing.Material{
				MaterialNumber: jsonRecipe.Materials[i].MaterialNumber,
				IngredientID:   jsonRecipe.Materials[i].IngredientID,
				Unit:           jsonRecipe.Materials[i].Unit,
				Amount:         jsonRecipe.Materials[i].Amount,
				Quality:        jsonRecipe.Materials[i].Quality,
				Note:           jsonRecipe.Materials[i].Note,
				Optional:       jsonRecipe.Materials[i].Optional,
			})
		}
		serviceRecipe.Materials = materials

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

	var materials []listing.Material
	for i := range jsonRecipe.Materials {
		materials = append(materials, listing.Material{
			MaterialNumber: jsonRecipe.Materials[i].MaterialNumber,
			IngredientID:   jsonRecipe.Materials[i].IngredientID,
			Unit:           jsonRecipe.Materials[i].Unit,
			Amount:         jsonRecipe.Materials[i].Amount,
			Quality:        jsonRecipe.Materials[i].Quality,
			Note:           jsonRecipe.Materials[i].Note,
			Optional:       jsonRecipe.Materials[i].Optional,
		})
	}
	serviceRecipe.Materials = materials

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
	if recipe.Materials != nil {
		var materials []Material
		for i := range recipe.Materials {
			materials = append(materials, Material{
				MaterialNumber: recipe.Materials[i].MaterialNumber,
				IngredientID:   recipe.Materials[i].IngredientID,
				Unit:           recipe.Materials[i].Unit,
				Amount:         recipe.Materials[i].Amount,
				Quality:        recipe.Materials[i].Quality,
				Note:           recipe.Materials[i].Note,
				Optional:       recipe.Materials[i].Optional,
			})
		}
		jsonRecipe.Materials = materials
	}

	if err := s.db.Write(CollectionRecipe, jsonRecipe.ID, jsonRecipe); err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteRecipe(id string) error {
	if err := s.db.Delete(CollectionRecipe, id); err != nil {
		return deleting.ErrNotFound
	}
	return nil
}

func (s *Storage) AddIngredient(ingredient adding.Ingredient) error {
	ingredientId, err := ksuid.NewRandom()
	if err != nil {
		return err
	}

	newIngredient := Ingredient{
		ID:          ingredientId.String(),
		Name:        ingredient.Name,
		Unit:        ingredient.Unit,
		MassInGrams: ingredient.MassInGrams,
		Substitutes: ingredient.Substitutes,
	}

	if err := s.db.Write(CollectionIngredient, newIngredient.ID, newIngredient); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetAllIngredients() []listing.Ingredient {
	list := []listing.Ingredient{}

	records, err := s.db.ReadAll(CollectionIngredient)
	if err != nil {
		return list
	}

	for _, r := range records {
		var jsonIngredient Ingredient
		var serviceIngredient listing.Ingredient

		if err := json.Unmarshal([]byte(r), &jsonIngredient); err != nil {
			return list
		}

		serviceIngredient.ID = jsonIngredient.ID
		serviceIngredient.Name = jsonIngredient.Name
		serviceIngredient.Unit = jsonIngredient.Unit
		serviceIngredient.MassInGrams = jsonIngredient.MassInGrams
		serviceIngredient.Substitutes = jsonIngredient.Substitutes

		list = append(list, serviceIngredient)
	}
	return list
}

func (s *Storage) GetIngredient(id string) (listing.Ingredient, error) {
	var serviceIngredient listing.Ingredient
	var jsonIngredient Ingredient

	err := s.db.Read(CollectionIngredient, id, &jsonIngredient)
	if err != nil {
		fmt.Println(err)
		return serviceIngredient, listing.ErrNotFound
	}

	serviceIngredient.ID = jsonIngredient.ID
	serviceIngredient.Name = jsonIngredient.Name
	serviceIngredient.Unit = jsonIngredient.Unit
	serviceIngredient.MassInGrams = jsonIngredient.MassInGrams
	serviceIngredient.Substitutes = jsonIngredient.Substitutes

	return serviceIngredient, nil
}

func (s *Storage) UpdateIngredient(id string, ingredient updating.Ingredient) error {
	var jsonIngredient Ingredient
	err := s.db.Read(CollectionIngredient, id, &jsonIngredient)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if ingredient.Name != "" {
		jsonIngredient.Name = ingredient.Name
	}
	if ingredient.Unit != "" {
		jsonIngredient.Unit = ingredient.Unit
	}
	if ingredient.MassInGrams != 0 {
		jsonIngredient.MassInGrams = ingredient.MassInGrams
	}
	if ingredient.Substitutes != "" {
		jsonIngredient.Substitutes = ingredient.Substitutes
	}

	if err := s.db.Write(CollectionIngredient, jsonIngredient.ID, jsonIngredient); err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteIngredient(id string) error {
	if err := s.db.Delete(CollectionIngredient, id); err != nil {
		return deleting.ErrNotFound
	}
	return nil
}

func (s *Storage) GetUnits() listing.UnitDict {
	var listingUnitDict listing.UnitDict
	var jsonUnitDict UnitDict
	err := s.db.Read(CollectionUnits, "12345", &jsonUnitDict)
	if err != nil {
		fmt.Println(err)
	}
	listingUnitDict.Dict = jsonUnitDict.Dict

	return listingUnitDict
}

func (s *Storage) AddUnit(unit string, units_per_cup float64) error {
	var jsonUnitDict UnitDict
	jsonUnitDict.Dict = make(map[string]float64)
	err := s.db.Read(CollectionUnits, "12345", &jsonUnitDict)
	if err != nil {
		fmt.Println(err)
	}
	jsonUnitDict.Dict[unit] = units_per_cup
	if err := s.db.Write(CollectionUnits, "12345", jsonUnitDict); err != nil {
		fmt.Println(err)
	}
	return nil
}
