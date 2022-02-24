package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mitchellss/recipe_rest/pkg/http/rest"
	"github.com/mitchellss/recipe_rest/pkg/service"
	"github.com/mitchellss/recipe_rest/pkg/storage/json"
)

func main() {
	var crudService service.Service
	repository, _ := json.NewStorage()

	crudService = service.NewService(repository)
	router := rest.Handler(crudService)

	fmt.Println("The recipe server live now at: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
