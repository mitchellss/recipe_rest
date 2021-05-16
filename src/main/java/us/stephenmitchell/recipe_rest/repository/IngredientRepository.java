package us.stephenmitchell.recipe_rest.repository;

import org.springframework.data.repository.CrudRepository;
import us.stephenmitchell.recipe_rest.model.Ingredient;

import java.util.List;

public interface IngredientRepository extends CrudRepository<Ingredient, Long> {
    List<Ingredient> findById(long id);
}
