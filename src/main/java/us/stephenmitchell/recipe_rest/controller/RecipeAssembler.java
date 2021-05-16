package us.stephenmitchell.recipe_rest.controller;

import org.springframework.stereotype.Component;
import us.stephenmitchell.recipe_rest.model.RecipeModel;

import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.*;

import org.springframework.hateoas.EntityModel;
import org.springframework.hateoas.server.RepresentationModelAssembler;
import org.springframework.stereotype.Component;

@Component
public class RecipeAssembler implements RepresentationModelAssembler<RecipeModel,EntityModel<RecipeModel>> {
    @Override
    public EntityModel<RecipeModel> toModel(RecipeModel recipe) {
        return EntityModel.of(recipe, //
                linkTo(methodOn(RecipeController.class).one(recipe.getId())).withSelfRel(),
                linkTo(methodOn(RecipeController.class).all()).withRel("recipe_list"));// )
    }
}
