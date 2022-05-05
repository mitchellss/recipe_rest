import { FunctionComponent, useState } from "react";

interface AddRecipePageProps {
}
 
const AddRecipePage : FunctionComponent<AddRecipePageProps> = () => {
    const [stepNums, setStepNums] = useState<number[]>();
    return (
        <form>
            <label htmlFor="recipe_name">Recipe Title:</label>
            <input type="text" id="recipe_name" name="recipe_name" />
            <br/>
            <label htmlFor="recipe_author">Author:</label>
            <input type="text" id="recipe_author" name="recipe_author" />
            <br/>
            <label htmlFor="active_time">Active Time:</label>
            <input type="number" id="active_time" name="active_time" step="1" pattern="\d+" />
            <label>minutes</label>
            <br/>
            <label htmlFor="total_time">Total Time:</label>
            <input type="number" id="total_time" name="total_time" step="1" pattern="\d+" />
            <label>minutes</label>
            <br/>
            <label htmlFor="serves_low">Serves (# of People):</label>
            <input type="number" id="fname" name="serves_low" step="1" pattern="\d+" />
            <label htmlFor="serves_high">to</label>
            <input type="number" id="serves_high" name="serves_high" step="1" pattern="\d+" />
            <h2>Ingredients:</h2>
            {
                <div>
                    <label>Ingredient:</label>
                    <input type="text" id="fname" name="fname"/>
                    <label>Amount:</label>
                    <input type="number" step="0.125" id="fname" name="fname"/>
                    <select name="unit_dropdown" id="1">
                        <option value="cup">Cup(s)</option>
                        <option value="tbsp">Tablespoon(s)</option>
                        <option value="tsp">Teaspoon(s)</option>
                    </select>
                    <br/>
                    <label>Adjectives (i.e. sliced, diced, packed, sifted):</label>
                    <input type="text" id="fname" name="fname"/>
                    <br/>
                    <label>Additional Notes:</label>
                    <input type="text" id="fname" name="fname"/>
                </div>
            }
            <input type="button" value="Add Ingredient"/>
            <br/>
            <h2>Steps:</h2>
            <ol>
                {
                stepNums?  
                    stepNums.map((item,index) => {
                        return(<li key={`step_${index}`}>
                            <label htmlFor={`step_${index}`}>Step Text:</label>
                            <input type="text" name={`step_${index}`}/>
                        </li>)
                    })
                    : <div></div>
                }
            </ol>
            <input type="button" value="Add Recipe Step"/>
        </form>
    );
}
 
export default AddRecipePage ;