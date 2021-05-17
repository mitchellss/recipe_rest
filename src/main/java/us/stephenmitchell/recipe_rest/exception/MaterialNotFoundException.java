package us.stephenmitchell.recipe_rest.exception;

public class MaterialNotFoundException extends RuntimeException {
    public MaterialNotFoundException(Long id) {
        super("Could not find material " + id);
    }
}
