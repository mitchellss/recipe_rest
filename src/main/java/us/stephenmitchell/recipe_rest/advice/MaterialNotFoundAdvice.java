package us.stephenmitchell.recipe_rest.advice;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import us.stephenmitchell.recipe_rest.exception.MaterialNotFoundException;

@ControllerAdvice
public class MaterialNotFoundAdvice {

    @ResponseBody
    @ExceptionHandler(MaterialNotFoundException.class)
    @ResponseStatus(HttpStatus.NOT_FOUND)
    String materialNotFoundHandler(MaterialNotFoundException ex) {
        return ex.getMessage();
    }
}
