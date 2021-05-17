package us.stephenmitchell.recipe_rest.exception;

public class StepNotFoundException extends RuntimeException {
    public StepNotFoundException(Long id) {
        super("Could not find step " + id);
    }
}
