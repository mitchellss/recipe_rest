package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellss/recipe_rest/pkg/service"
)

func Handler(s service.Service) http.Handler {
	router := httprouter.New()
	router.GET("/recipes", getRecipes(s))
	router.POST("/add", addRecipe(s))
	return router
}

func getRecipes(service service.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := service.GetAllRecipes()
		json.NewEncoder(w).Encode(list)
	}
}

func addRecipe(crudService service.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var newRecipe service.Recipe

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
