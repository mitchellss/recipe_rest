package us.stephenmitchell.recipe_rest.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import us.stephenmitchell.recipe_rest.repository.RecipeRepository;
import us.stephenmitchell.recipe_rest.model.RecipeModel;

@RestController
@RequestMapping("/api")
public class RecipeController {

    @Autowired
    RecipeRepository recipeRepository;

    @GetMapping("/get_recipe")
    public Iterable<RecipeModel> getRecipe() {
        return recipeRepository.findAll();
    }

    @PostMapping("/post_recipe")
    public String postRecipe(@RequestBody RecipeModel recipe) {
        recipeRepository.save(recipe);
        return recipe.toString();
    }

}
