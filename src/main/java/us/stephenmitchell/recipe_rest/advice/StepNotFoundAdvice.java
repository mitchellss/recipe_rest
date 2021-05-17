package us.stephenmitchell.recipe_rest.advice;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import us.stephenmitchell.recipe_rest.exception.StepNotFoundException;

@ControllerAdvice
public class StepNotFoundAdvice {

    @ResponseBody
    @ExceptionHandler(StepNotFoundException.class)
    @ResponseStatus(HttpStatus.NOT_FOUND)
    String stepNotFoundHandler(StepNotFoundException ex) {
        return ex.getMessage();
    }
}
