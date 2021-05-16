package us.stephenmitchell.recipe_rest.exception;

public class IngredientNotFoundException extends RuntimeException {
    public IngredientNotFoundException(Long id) {
        super("Could not find ingredient" + id);
    }
}
