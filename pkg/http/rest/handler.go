package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellss/recipe_rest/pkg/adding"
	"github.com/mitchellss/recipe_rest/pkg/listing"
)

func Handler(a adding.Service, l listing.Service) http.Handler {
	router := httprouter.New()
	router.GET("/recipes", getRecipes(l))
	router.GET("/recipe/:id", getRecipe(l))
	router.POST("/add", addRecipe(a))
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
		list := crudService.GetAllRecipes()
		json.NewEncoder(w).Encode(list)
	}
}
