package us.stephenmitchell.recipe_rest.assembler;

import org.springframework.hateoas.EntityModel;
import org.springframework.hateoas.server.RepresentationModelAssembler;
import org.springframework.hateoas.server.mvc.WebMvcLinkBuilder;
import org.springframework.stereotype.Component;
import us.stephenmitchell.recipe_rest.controller.IngredientController;
import us.stephenmitchell.recipe_rest.controller.RecipeController;
import us.stephenmitchell.recipe_rest.model.Ingredient;
import us.stephenmitchell.recipe_rest.model.Recipe;
import us.stephenmitchell.recipe_rest.repository.IngredientRepository;

import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.linkTo;
import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.methodOn;

@Component
public class IngredientAssembler implements RepresentationModelAssembler<Ingredient,EntityModel<Ingredient>> {
    @Override
    public EntityModel<Ingredient> toModel(Ingredient ingredient) {
        return EntityModel.of(ingredient, //
                WebMvcLinkBuilder.linkTo(methodOn(IngredientController.class).one(ingredient.getId())).withSelfRel(),
                linkTo(methodOn(IngredientController.class).all()).withRel("ingredient_list"));// )
    }
}
