package us.stephenmitchell.recipe_rest.assembler;

import org.springframework.hateoas.EntityModel;
import org.springframework.hateoas.server.RepresentationModelAssembler;
import org.springframework.hateoas.server.mvc.WebMvcLinkBuilder;
import org.springframework.stereotype.Component;
import us.stephenmitchell.recipe_rest.controller.RecipeController;
import us.stephenmitchell.recipe_rest.controller.StepController;
import us.stephenmitchell.recipe_rest.model.Step;

import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.linkTo;
import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.methodOn;

@Component
public class StepAssembler implements RepresentationModelAssembler<Step,EntityModel<Step>> {
    @Override
    public EntityModel<Step> toModel(Step step) {
        return EntityModel.of(step, //
                WebMvcLinkBuilder.linkTo(methodOn(StepController.class).one(step.getId())).withSelfRel(),
                linkTo(methodOn(RecipeController.class).all()).withRel("step_list"),
                linkTo(methodOn(RecipeController.class).all()).withRel("recipe"));
    }
}
