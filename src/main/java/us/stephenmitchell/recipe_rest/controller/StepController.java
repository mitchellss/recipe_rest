package us.stephenmitchell.recipe_rest.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.hateoas.EntityModel;
import org.springframework.web.bind.annotation.*;
import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.*;

import us.stephenmitchell.recipe_rest.model.StepModel;
import us.stephenmitchell.recipe_rest.repository.StepRepository;

@RestController
@RequestMapping("/api")
public class StepController {

    @Autowired
    StepRepository stepRepository;

    @GetMapping("/get_step")
    public Iterable<StepModel> all() {
        return stepRepository.findAll();
    }

    @PostMapping("/post_step")
    public String postStep(@RequestBody StepModel step) {
        stepRepository.save(step);
        return step.toString();
    }

    @GetMapping("/get_step/{id}")
    public EntityModel<StepModel> one(@PathVariable Long id) {
        StepModel step = stepRepository.findById(id)
                .orElseThrow(() -> new StepNotFoundException(id));

        return EntityModel.of(step,
                linkTo(methodOn(StepController.class).one(id)).withSelfRel(),
                linkTo(methodOn(StepController.class).all()).withRel("step_list"),
                linkTo(methodOn(RecipeController.class).one(step.getRecipe().getId())).withRel("recipe"));
    }
}
