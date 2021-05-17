package us.stephenmitchell.recipe_rest.advice;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import us.stephenmitchell.recipe_rest.exception.IngredientNotFoundException;
import us.stephenmitchell.recipe_rest.exception.RecipeNotFoundException;
import us.stephenmitchell.recipe_rest.model.Ingredient;

@ControllerAdvice
public class IngredientNotFoundAdvice {

    @ResponseBody
    @ExceptionHandler(IngredientNotFoundException.class)
    @ResponseStatus(HttpStatus.NOT_FOUND)
    String ingredientNotFoundHandler(IngredientNotFoundException ex) {
        return ex.getMessage();
    }
}
