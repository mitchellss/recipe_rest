package us.stephenmitchell.recipe_rest.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.hateoas.CollectionModel;
import org.springframework.hateoas.EntityModel;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.*;

import us.stephenmitchell.recipe_rest.repository.RecipeRepository;
import us.stephenmitchell.recipe_rest.model.RecipeModel;

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

    // @GetMapping("/get_recipe")
    // public Iterable<RecipeModel> all() {
    //     return recipeRepository.findAll();
    // }

    @GetMapping("/get_recipe")
    public CollectionModel<EntityModel<RecipeModel>> all() {
        List<EntityModel<RecipeModel>> recipes = StreamSupport
                .stream(recipeRepository.findAll().spliterator(), false)
                .map(recipeAssembler::toModel)
                .collect(Collectors.toList());
        return CollectionModel.of(recipes,
                linkTo(methodOn(RecipeController.class).all()).withSelfRel());
    }

    @PostMapping("/post_recipe")
    public String postRecipe(@RequestBody RecipeModel recipe) {
        recipeRepository.save(recipe);
        return recipe.toString();
    }

    @GetMapping("/get_recipe/{id}")
    public EntityModel<RecipeModel> one(@PathVariable Long id) {
        RecipeModel recipe = recipeRepository.findById(id)
                .orElseThrow(() -> new RecipeNotFoundException(id));

        return recipeAssembler.toModel(recipe);
    }

}
