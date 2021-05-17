package us.stephenmitchell.recipe_rest.repository;

import org.springframework.data.repository.CrudRepository;
import us.stephenmitchell.recipe_rest.model.Material;

import java.util.List;

public interface MaterialRepository extends CrudRepository<Material, Long> {
    List<Material> findById(long id);
}
