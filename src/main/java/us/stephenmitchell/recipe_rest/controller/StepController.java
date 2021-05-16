package us.stephenmitchell.recipe_rest.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import us.stephenmitchell.recipe_rest.model.StepModel;
import us.stephenmitchell.recipe_rest.repository.StepRepository;

@RestController
@RequestMapping("/api")
public class StepController {

    @Autowired
    StepRepository stepRepository;

    @GetMapping("/get_step")
    public Iterable<StepModel> getStep() { return stepRepository.findAll(); }

    @PostMapping("/post_step")
    public String postStep(@RequestBody StepModel step) {
        stepRepository.save(step);
        return step.toString();
    }

    @GetMapping("/get_step/{id}")
    public StepModel one(@PathVariable Long id) {
        return stepRepository.findById(id)
                .orElseThrow(() -> new StepNotFoundException(id));
    }
}
