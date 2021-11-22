package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mitchellss/recipe_rest"
	"gopkg.in/yaml.v3"
)

func main() {
	file := flag.String("file", "recipe.yaml", "The YAML file that contains the recipe")
	//port := flag.Int("port", 3000, "The port to host the application on")
	flag.Parse()

	fmt.Printf("The recipe being displayed is %s\n", *file)

	recipeFile, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	var recipe recipe_rest.Recipe

	err = yaml.Unmarshal(recipeFile, &recipe)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", recipe)

	// handler := recipe_rest.NewHandler(recipe)

	// fmt.Printf("Starting server on port %d\n", *port)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
