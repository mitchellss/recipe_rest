package us.stephenmitchell.recipe_rest.repository;

import org.springframework.data.repository.CrudRepository;
import us.stephenmitchell.recipe_rest.model.StepModel;

import java.util.List;

public interface StepRepository extends CrudRepository<StepModel, Long> {
    List<StepModel> findById(long id);
}
