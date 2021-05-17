package us.stephenmitchell.recipe_rest.repository;

import org.springframework.data.repository.CrudRepository;
import us.stephenmitchell.recipe_rest.model.Step;

import java.util.List;

public interface StepRepository extends CrudRepository<Step, Long> {
    List<Step> findById(long id);
}
