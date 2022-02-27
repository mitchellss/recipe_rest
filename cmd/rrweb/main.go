package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mitchellss/recipe_rest/pkg/adding"
	"github.com/mitchellss/recipe_rest/pkg/http/rest"
	"github.com/mitchellss/recipe_rest/pkg/listing"
	"github.com/mitchellss/recipe_rest/pkg/storage/json"
)

func main() {
	var adder adding.Service
	var lister listing.Service

	repository, _ := json.NewStorage()

	adder = adding.NewService(repository)
	lister = listing.NewService(repository)

	router := rest.Handler(adder, lister)

	fmt.Println("The recipe server live now at: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
