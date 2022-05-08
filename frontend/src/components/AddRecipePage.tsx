import axios from "axios";
import React, { FunctionComponent, useState } from "react";
import * as uuid from "uuid";
import { HOST, PORT } from "./constants";

interface RecipeStep {
    stepId: string
}

interface RecipeIngredient {
    ingredientId: string,
}

interface AddRecipePageProps {
}

const AddRecipePage: FunctionComponent<AddRecipePageProps> = () => {
    const [steps, setSteps] = useState<RecipeStep[]>();
    const [ingredients, setIngredients] = useState<RecipeIngredient[]>();

    const clickAddIngredient = () => {
        const ingredient_id = uuid.v4();
        const new_ingredient: RecipeIngredient = {
            ingredientId: ingredient_id
        }

        ingredients ?
            setIngredients([...ingredients, new_ingredient])
            : setIngredients([new_ingredient])
    }

    const clickAddStep = () => {
        const step_id = uuid.v4();
        const new_step: RecipeStep = {
            stepId: step_id
        }

        steps ?
            setSteps([...steps, new_step])
            : setSteps([new_step])
    }

    const clickDeleteStep = (stepId: string) => {
        const newSteps = steps && steps.filter((element, i) => element.stepId !== stepId)
        setSteps(newSteps)
    }

    const clickDeleteIngredient = (stepId: string) => {
        const newIngredients = ingredients && ingredients.filter((element, i) => element.ingredientId !== stepId)
        setIngredients(newIngredients)
    }

    const populateSteps = (): JSX.Element[] => {
        const recipeSteps: JSX.Element[] = [];
        steps ?
            steps.map((item, index) => {
                recipeSteps.push(
                    <li key={`step_${item.stepId}`} className="InputLine">
                        <div>
                            <label htmlFor={`step_text_${item.stepId}`}>Step Text:</label>
                        </div>
                        <div>
                            <input type="text" name={`step_text_${item.stepId}`} />
                            <input type="button" value="Delete" className="DeleteButton" onClick={() => clickDeleteStep(item.stepId)} />
                        </div>
                    </li>
                )
            })
            : recipeSteps.push(<div></div>);
        return recipeSteps;
    }

    const populateIngredients = (): JSX.Element[] => {
        const recipeIngredients: JSX.Element[] = [];
        ingredients ?
            ingredients.map((item, index) => {
                recipeIngredients.push(
                    <li key={`${item.ingredientId}`} className="IngredientBlock">
                        <div className="InputLine InputLineIngredient">
                            <label>Ingredient:</label>
                            <input type="text" name={`ingredient_name_${item.ingredientId}`} />
                            <label>Amount:</label>
                            <input type="number" step="0.125" name={`ingredient_amount_${item.ingredientId}`} />
                            <select name={`ingredient_unit_${item.ingredientId}`} id="1">
                                <option value="cup">Cup(s)</option>
                                <option value="tablespoon">Tablespoon(s)</option>
                                <option value="teaspoon">Teaspoon(s)</option>
                            </select>
                        </div>
                        <div className="InputLine InputLineIngredient">
                            <label>Adjectives (i.e. sliced, diced, packed, sifted):</label>
                            <input type="text" name={`ingredient_adj_${item.ingredientId}`} />
                        </div>
                        <div className="InputLine InputLineIngredient">
                            <label>Additional Notes:</label>
                            <input type="text" name={`ingredient_note_${item.ingredientId}`} />
                            <input type="button" value="Delete" className="DeleteButton" onClick={() => { clickDeleteIngredient(item.ingredientId) }} />
                        </div>
                    </li>
                )
            })
            : recipeIngredients.push(<div></div>);
        return recipeIngredients;
    }

    const submitRecipe = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        var addRecipeSteps: object[] = [];
        steps?.map((item, index) => {
            const step_text = (event.currentTarget.elements.namedItem(`step_text_${item.stepId}`) as HTMLInputElement).value
            addRecipeSteps.push({
                "step_num": index + 1,
                "text" : step_text,
            })
        });

        var addRecipeIngredients: object[] = [];
        ingredients?.map((item, index) => {
            const ingredient_id = "test123";
            const ingredient_amount = (event.currentTarget.elements.namedItem(`ingredient_amount_${item.ingredientId}`) as HTMLInputElement).value
            const ingredient_unit = (event.currentTarget.elements.namedItem(`ingredient_unit_${item.ingredientId}`) as HTMLInputElement).value
            const ingredient_adj = (event.currentTarget.elements.namedItem(`ingredient_adj_${item.ingredientId}`) as HTMLInputElement).value
            const ingredient_note = (event.currentTarget.elements.namedItem(`ingredient_note_${item.ingredientId}`) as HTMLInputElement).value
            addRecipeIngredients.push({
                "material_num": index + 1,
                "ingredient_id": ingredient_id,
                "unit": ingredient_unit,
                "amount": +ingredient_amount,
                "quality": ingredient_adj,
                "note": ingredient_note
            })
        });

        const data = {
            "title": event.currentTarget.recipe_name.value,
            "author" : event.currentTarget.recipe_author.value,
            "active_time" : +event.currentTarget.active_time.value,
            "total_time" : +event.currentTarget.total_time.value,
            "serves_high" : +event.currentTarget.serves_low.value,
            "serves_low" : +event.currentTarget.serves_high.value,
            "steps" : addRecipeSteps,
            "materials" : addRecipeIngredients
        }

        axios.post(`${HOST}:${PORT}/api/recipe/`, data).then((res) => {
            console.log(res)
        })
    }

    const cancelSubmitRecipe = (): void => {

    }

    return (
        <form onSubmit={submitRecipe}>
            <h1>Add a Recipe:</h1>
            <div className="InputLine">
                <label htmlFor="recipe_name" className="InputLabel">Recipe Title:</label>
                <div className="InputContent">
                    <input type="text" id="recipe_name" name="recipe_name" />
                </div>
            </div>
            <div className="InputLine">
                <label htmlFor="recipe_author" className="InputLabel">Author:</label>
                <div className="InputContent">
                    <input type="text" id="recipe_author" name="recipe_author" />
                </div>
            </div>
            <div className="InputLine">
                <label htmlFor="active_time" className="InputLabel">Active Time:</label>
                <div className="InputContent">
                    <input type="number" id="active_time" name="active_time" step="1" pattern="\d+" />
                    <label>minutes</label>
                </div>

            </div>
            <div className="InputLine">
                <label htmlFor="total_time" className="InputLabel">Total Time:</label>
                <div className="InputContent">
                    <input type="number" id="total_time" name="total_time" step="1" pattern="\d+" />
                    <label>minutes</label>
                </div>

            </div>
            <div className="InputLine">
                <label htmlFor="serves_low">Serves (# of People):</label>
                <div className="InputContent">
                    <input type="number" name="serves_low" step="1" pattern="\d+" />
                    <label htmlFor="serves_high">to</label>
                    <input type="number" id="serves_high" name="serves_high" step="1" pattern="\d+" />
                </div>

            </div>
            <h2>Ingredients:</h2>
            <ul className="NoULStyle">
                {
                    populateIngredients()
                }
            </ul>
            <input type="button" value="Add Ingredient" onClick={clickAddIngredient} />
            <br />
            <h2>Steps:</h2>
            <ol className="FixedWidthOL">
                {
                    populateSteps()
                }
            </ol>
            <input type="button" value="Add Recipe Step" onClick={clickAddStep} />
                <div>
                <input type="submit" value="Add Recipe" className="SubmitRecipeButton"/>
                <input type="button" value="Cancel" className="SubmitRecipeButton" onClick={cancelSubmitRecipe}/>
                </div>
        </form>
    );
}

export default AddRecipePage;