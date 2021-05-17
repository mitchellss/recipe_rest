package us.stephenmitchell.recipe_rest;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.CommandLineRunner;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import us.stephenmitchell.recipe_rest.model.Ingredient;
import us.stephenmitchell.recipe_rest.model.Material;
import us.stephenmitchell.recipe_rest.model.Recipe;
import us.stephenmitchell.recipe_rest.model.Step;
import us.stephenmitchell.recipe_rest.repository.IngredientRepository;
import us.stephenmitchell.recipe_rest.repository.MaterialRepository;
import us.stephenmitchell.recipe_rest.repository.RecipeRepository;
import us.stephenmitchell.recipe_rest.repository.StepRepository;

@Configuration
public class LoadDatabase {

    private static final Logger log = LoggerFactory.getLogger(LoadDatabase.class);

    @Bean
    CommandLineRunner initDatabase(IngredientRepository ingredientRepository, MaterialRepository materialRepository,
                                   RecipeRepository recipeRepository, StepRepository stepRepository) {
        Ingredient testIngredient = new Ingredient("Test1", "test2", 5L);
        Recipe testRecipe = new Recipe("TestRecipe", "1970-01-01 01:00:00");
        Material testMaterial = new Material(testRecipe, 1L, "1 Cup",
                testIngredient, "testNote");
        Step testStep = new Step(1L, "testStep", testRecipe);
        return args -> {
            log.info("Preloading " + ingredientRepository.save(testIngredient));
            log.info("Preloading " + recipeRepository.save(testRecipe));
            log.info("Preloading " + materialRepository.save(testMaterial));
            log.info("Preloading " + stepRepository.save(testStep));
        };
    }
}
