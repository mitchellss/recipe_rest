package us.stephenmitchell.recipe_rest.assembler;

import org.springframework.hateoas.EntityModel;
import org.springframework.hateoas.server.RepresentationModelAssembler;
import org.springframework.hateoas.server.mvc.WebMvcLinkBuilder;
import org.springframework.stereotype.Component;
import us.stephenmitchell.recipe_rest.controller.IngredientController;
import us.stephenmitchell.recipe_rest.controller.MaterialController;
import us.stephenmitchell.recipe_rest.controller.RecipeController;
import us.stephenmitchell.recipe_rest.model.Ingredient;
import us.stephenmitchell.recipe_rest.model.Material;
import us.stephenmitchell.recipe_rest.model.Recipe;

import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.linkTo;
import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.methodOn;

@Component
public class MaterialAssembler implements RepresentationModelAssembler<Material,EntityModel<Material>> {
    @Override
    public EntityModel<Material> toModel(Material material) {
        return EntityModel.of(material, //
                WebMvcLinkBuilder.linkTo(methodOn(MaterialController.class).one(material.getId())).withSelfRel(),
                linkTo(methodOn(MaterialController.class).all()).withRel("material_list"),
                linkTo(methodOn(IngredientController.class).one(material.getIngredient().getId())).withRel("ingredient"),
                linkTo(methodOn(RecipeController.class).one(material.getRecipe().getId())).withRel("recipe"));
    }

}
