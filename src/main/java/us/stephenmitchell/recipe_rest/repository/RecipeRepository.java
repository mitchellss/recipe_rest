package us.stephenmitchell.recipe_rest.repository;

import org.springframework.data.repository.CrudRepository;
import us.stephenmitchell.recipe_rest.model.RecipeModel;

import java.util.List;

public interface RecipeRepository extends CrudRepository<RecipeModel, Long> {
    List<RecipeModel> findById(long id);
}
