package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellss/recipe_rest/pkg/adding"
	"github.com/mitchellss/recipe_rest/pkg/deleting"
	"github.com/mitchellss/recipe_rest/pkg/listing"
	"github.com/mitchellss/recipe_rest/pkg/updating"
)

func Handler(c adding.Service, r listing.Service, u updating.Service, d deleting.Service) http.Handler {
	router := httprouter.New()
	router.POST("/api/recipe", addRecipe(c))
	router.POST("/api/ingredient", addIngredient(c))
	router.POST("/api/unit", addUnit(c))
	router.GET("/api/recipe", getRecipes(r))
	router.GET("/api/ingredient", getIngredients(r))
	router.GET("/api/recipe/:id", getRecipe(r))
	router.GET("/api/ingredient/:id", getIngredient(r))
	router.GET("/api/unit", getUnits(r))
	router.PUT("/api/recipe/:id", updateRecipe(u))
	router.PUT("/api/ingredient/:id", updateIngredient(u))
	router.DELETE("/api/recipe/:id", deleteRecipe(d))
	router.DELETE("/api/ingredient/:id", deleteIngredient(d))
	return router
}

func addRecipe(crudService adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var newRecipe adding.Recipe

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newRecipe)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		crudService.AddRecipe(newRecipe)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New recipe added.")
	}
}

func addIngredient(crudService adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var newIngredient adding.Ingredient

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newIngredient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		crudService.AddIngredient(newIngredient)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New ingredient added.")

	}
}

func addUnit(crudService adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var newUnit adding.UnitDict

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newUnit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		crudService.AddUnit(newUnit)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New unit added.")
	}
}

func getRecipe(crudService listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		recipe, err := crudService.GetRecipe(p.ByName("id"))
		if err == listing.ErrNotFound {
			http.Error(w, "The recipe you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recipe)
	}
}

func getRecipes(crudService listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		query := r.URL.Query()
		var list []listing.Recipe
		val, ok := query["id"]
		if ok {
			for _, val := range val {
				recipe, err := crudService.GetRecipe(val)
				if err == listing.ErrNotFound {
					http.Error(w, "The recipe you requested does not exist.", http.StatusNotFound)
					return
				}
				list = append(list, recipe)
			}
		} else {
			list = crudService.GetAllRecipes()
		}

		json.NewEncoder(w).Encode(list)
	}
}

func getUnits(crudService listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := crudService.GetUnits()
		json.NewEncoder(w).Encode(list)
	}
}

func updateRecipe(crudService updating.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var newRecipe updating.Recipe

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newRecipe)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = crudService.UpdateRecipe(p.ByName("id"), newRecipe)
		if err == updating.ErrNotFound {
			http.Error(w, "The recipe you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Recipe updated.")
	}

}

func deleteRecipe(crudService deleting.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := crudService.DeleteRecipe(p.ByName("id"))
		if err == deleting.ErrNotFound {
			http.Error(w, "The recipe you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Recipe deleted.")

	}
}

func getIngredients(crudService listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		query := r.URL.Query()
		var list []listing.Ingredient
		val, ok := query["id"]
		if ok {
			for _, val := range val {
				ingredient, err := crudService.GetIngredient(val)
				if err == listing.ErrNotFound {
					http.Error(w, "The ingredient you requested does not exist.", http.StatusNotFound)
					return
				}
				list = append(list, ingredient)
			}
		} else {
			list = crudService.GetAllIngredients()
		}
		json.NewEncoder(w).Encode(list)
	}
}

func getIngredient(crudService listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ingredient, err := crudService.GetIngredient(p.ByName("id"))
		if err == listing.ErrNotFound {
			http.Error(w, "The ingredient you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ingredient)
	}
}

func updateIngredient(crudService updating.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var newIngredient updating.Ingredient

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newIngredient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = crudService.UpdateIngredient(p.ByName("id"), newIngredient)
		if err == updating.ErrNotFound {
			http.Error(w, "The ingredient you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Ingredient updated.")
	}

}

func deleteIngredient(crudService deleting.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := crudService.DeleteIngredient(p.ByName("id"))
		if err == deleting.ErrNotFound {
			http.Error(w, "The ingredient you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Ingredient deleted.")

	}
}
