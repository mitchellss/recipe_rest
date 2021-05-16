package us.stephenmitchell.recipe_rest.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.hateoas.CollectionModel;
import org.springframework.hateoas.EntityModel;
import org.springframework.web.bind.annotation.*;
import us.stephenmitchell.recipe_rest.assembler.IngredientAssembler;
import us.stephenmitchell.recipe_rest.exception.IngredientNotFoundException;
import us.stephenmitchell.recipe_rest.exception.RecipeNotFoundException;
import us.stephenmitchell.recipe_rest.model.Ingredient;
import us.stephenmitchell.recipe_rest.model.Recipe;
import us.stephenmitchell.recipe_rest.model.Step;
import us.stephenmitchell.recipe_rest.repository.IngredientRepository;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.StreamSupport;

import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.linkTo;
import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.methodOn;

@RestController
@RequestMapping("/api")
public class IngredientController {

    @Autowired
    IngredientRepository ingredientRepository;

    IngredientAssembler ingredientAssembler;

    IngredientController(IngredientRepository ingredientRepository, IngredientAssembler ingredientAssembler) {
        this.ingredientRepository = ingredientRepository;
        this.ingredientAssembler = ingredientAssembler;
    }


    @GetMapping("/get_ingredient")
    public CollectionModel<EntityModel<Ingredient>> all() {
        List<EntityModel<Ingredient>> recipes = StreamSupport
                .stream(ingredientRepository.findAll().spliterator(), false)
                .map(ingredientAssembler::toModel)
                .collect(Collectors.toList());
        return CollectionModel.of(recipes,
                linkTo(methodOn(IngredientController.class).all()).withSelfRel());
    }

    @PostMapping("/post_ingredient")
    public String postIngredient(@RequestBody Ingredient ingredient) {
        ingredientRepository.save(ingredient);
        return ingredient.toString();
    }

    @GetMapping("/get_ingredient/{id}")
    public EntityModel<Ingredient> one(@PathVariable Long id) {
        Ingredient ingredient= ingredientRepository.findById(id)
                .orElseThrow(() -> new IngredientNotFoundException(id));

        return ingredientAssembler.toModel(ingredient);
    }
}
