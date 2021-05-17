package us.stephenmitchell.recipe_rest.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.hateoas.CollectionModel;
import org.springframework.hateoas.EntityModel;
import org.springframework.web.bind.annotation.*;
import static org.springframework.hateoas.server.mvc.WebMvcLinkBuilder.*;

import us.stephenmitchell.recipe_rest.assembler.StepAssembler;
import us.stephenmitchell.recipe_rest.exception.StepNotFoundException;
import us.stephenmitchell.recipe_rest.exception.StepNotFoundException;
import us.stephenmitchell.recipe_rest.model.Step;
import us.stephenmitchell.recipe_rest.repository.StepRepository;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.StreamSupport;

@RestController
@RequestMapping("/api")
public class StepController {

    @Autowired
    StepRepository stepRepository;
    StepAssembler stepAssembler;

    StepController(StepRepository stepRepository, StepAssembler stepAssembler) {
        this.stepRepository = stepRepository;
        this.stepAssembler = stepAssembler;
    }

    @GetMapping("/step")
    public CollectionModel<EntityModel<Step>> all() {
        List<EntityModel<Step>> steps = StreamSupport
                .stream(stepRepository.findAll().spliterator(), false)
                .map(stepAssembler::toModel)
                .collect(Collectors.toList());
        return CollectionModel.of(steps,
                linkTo(methodOn(StepController.class).all()).withSelfRel());

    }

    @PostMapping("/step")
    public String postStep(@RequestBody Step step) {
        stepRepository.save(step);
        return step.toString();
    }

    @GetMapping("/step/{id}")
    public EntityModel<Step> one(@PathVariable Long id) {
        Step step = stepRepository.findById(id)
                .orElseThrow(() -> new StepNotFoundException(id));

        return stepAssembler.toModel(step);
    }

    @PutMapping("/step/{id}")
    public EntityModel<Step> replaceStep(@RequestBody Step newStep, @PathVariable Long id) {
        Step step = stepRepository.findById(id)
                .orElseThrow(() -> new StepNotFoundException(id));

        step.setStep_number(newStep.getStep_number());
        step.setText(newStep.getText());
        step.setRecipe(newStep.getRecipe());
        stepRepository.save(step);
        return stepAssembler.toModel(step);
    }

    @DeleteMapping("/step/{id}")
    public void deleteStep(@PathVariable Long id) {
        stepRepository.deleteById(id);
    }
}
