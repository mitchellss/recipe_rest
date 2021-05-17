package us.stephenmitchell.recipe_rest.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.hateoas.CollectionModel;
import org.springframework.hateoas.EntityModel;
import org.springframework.web.bind.annotation.*;
import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.*;

import us.stephenmitchell.recipe_rest.assembler.RecipeAssembler;
import us.stephenmitchell.recipe_rest.exception.RecipeNotFoundException;
import us.stephenmitchell.recipe_rest.exception.RecipeNotFoundException;
import us.stephenmitchell.recipe_rest.model.Recipe;
import us.stephenmitchell.recipe_rest.repository.RecipeRepository;
import us.stephenmitchell.recipe_rest.model.Recipe;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.StreamSupport;

@RestController
@RequestMapping("/api")
public class RecipeController {

    @Autowired
    RecipeRepository recipeRepository;

    private final RecipeAssembler recipeAssembler;

    RecipeController(RecipeRepository recipeRepository, RecipeAssembler recipeAssembler) {
        this.recipeRepository = recipeRepository;
        this.recipeAssembler = recipeAssembler;
    }

    @GetMapping("/recipe")
    public CollectionModel<EntityModel<Recipe>> all() {
        List<EntityModel<Recipe>> recipes = StreamSupport
                .stream(recipeRepository.findAll().spliterator(), false)
                .map(recipeAssembler::toModel)
                .collect(Collectors.toList());
        return CollectionModel.of(recipes,
                linkTo(methodOn(RecipeController.class).all()).withSelfRel());
    }

    @PostMapping("/recipe")
    public String postRecipe(@RequestBody Recipe recipe) {
        recipeRepository.save(recipe);
        return recipe.toString();
    }

    @GetMapping("/recipe/{id}")
    public EntityModel<Recipe> one(@PathVariable Long id) {
        Recipe recipe = recipeRepository.findById(id)
                .orElseThrow(() -> new RecipeNotFoundException(id));

        return recipeAssembler.toModel(recipe);
    }

    @PutMapping("/recipe/{id}")
    public EntityModel<Recipe> replaceRecipe(@RequestBody Recipe newRecipe, @PathVariable Long id) {
        Recipe recipe = recipeRepository.findById(id)
                .orElseThrow(() -> new RecipeNotFoundException(id));

        recipe.setTitle(newRecipe.getTitle());
        recipe.setDatetime(newRecipe.getDatetime());
        recipeRepository.save(recipe);
        return recipeAssembler.toModel(recipe);
    }

    @DeleteMapping("/recipe/{id}")
    public void deleteRecipe(@PathVariable Long id) {
        recipeRepository.deleteById(id);
    }
}
