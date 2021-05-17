package us.stephenmitchell.recipe_rest.repository;

import org.springframework.data.repository.CrudRepository;
import us.stephenmitchell.recipe_rest.model.Recipe;

import java.util.List;

public interface RecipeRepository extends CrudRepository<Recipe, Long> {
    List<Recipe> findById(long id);
}
