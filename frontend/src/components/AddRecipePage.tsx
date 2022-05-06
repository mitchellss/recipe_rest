import { FunctionComponent, useState } from "react";
import * as uuid from "uuid";

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
                            <label htmlFor={`step_${item.stepId}`}>Step Text:</label>
                        </div>
                        <div>
                            <input type="text" name={`step_${item.stepId}`} />
                            <input type="button" value="delete" onClick={() => clickDeleteStep(item.stepId)} />
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
                            <input type="text" id="fname" name="fname" />
                            <label>Amount:</label>
                            <input type="number" step="0.125" id="fname" name="fname" />
                            <select name="unit_dropdown" id="1">
                                <option value="cup">Cup(s)</option>
                                <option value="tbsp">Tablespoon(s)</option>
                                <option value="tsp">Teaspoon(s)</option>
                            </select>
                        </div>
                        <div className="InputLine InputLineIngredient">
                            <label>Adjectives (i.e. sliced, diced, packed, sifted):</label>
                            <input type="text" id="fname" name="fname" />
                        </div>
                        <div className="InputLine InputLineIngredient">
                            <label>Additional Notes:</label>
                            <input type="text" id="fname" name="fname" />
                            <input type="button" id="fname" name="fname" value="Delete" className="DeleteButton" onClick={() => { clickDeleteIngredient(item.ingredientId) }} />
                        </div>
                    </li>
                )
            })
            : recipeIngredients.push(<div></div>);
        return recipeIngredients;
    }

    return (
        <form>
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
                    <input type="number" id="fname" name="serves_low" step="1" pattern="\d+" />
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
        </form>
    );
}

export default AddRecipePage;