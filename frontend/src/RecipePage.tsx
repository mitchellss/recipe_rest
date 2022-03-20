import axios from "axios";
import { useState, useEffect } from "react";
import {useParams} from "react-router-dom";
import { HOST, PORT } from "./constants";


function RecipePage() {
    const params = useParams()
    const [recipeData, setRecipeData] = useState<Recipe>();
    const [ingredientData, setIngredientData] = useState<Ingredient[]>();
    const [unitData, setUnitData] = useState<Map<string, number>>();

    const getAllData = () => {
        axios.get(`${HOST}:${PORT}/api/recipe/${params.recipeId}`)
            .then(res => {
                // console.log(res.data)
                setRecipeData(res.data);
                getIngredientData(res.data.materials);
                getUnitData();
            })
            .catch((error) => {
                console.log(error);
            });
    }

    useEffect(() => {
        getAllData();
    }, []);

    const getIngredientData = (materials: Material[]) => {
        var queryString = "?"
        materials.map( (material) => {
            queryString = queryString + "&id=" + material.ingredient_id
        })
        axios.get(`${HOST}:${PORT}/api/ingredient${queryString}`)
            .then(res => {
                // console.log(res.data)
                setIngredientData(res.data)
            })
            .catch((error) => {
                console.log(error);
            });
        // console.log(queryString)
    }

    const getUnitData = () => {
        axios.get(`${HOST}:${PORT}/api/unit`)
        .then(res => {
            var a: Map<string, number> = new Map();
            for (const item in res.data.Dict) {
                a.set(item, res.data.Dict[item])
                // console.log(item)
                // console.log(res.data.Dict[item])
            }
            // console.log(a)
            // console.log(res.data.Dict)
            setUnitData(a)
        })
        .catch((error) => {
            console.log(error);
        });

    }

    return ( 
    <div>
        {
            recipeData ? 
                <div>
                    <div className="Title RecipeTitle">{recipeData.title}</div>
                    <div className="Title RecipeAuthor">{recipeData.author}</div>
                    <div className="Title RecipeCreated">{recipeData.created}</div>
                    <br />
                    <div className="RecipePageBox IngredientsBox">
                        <div className="RecipePageTitle IngredientsTitle">Ingredients</div>
                        {recipeData.materials.map(mat => {
                            // console.log(mat);
                            
                            var mass = -1;

                            // ingredientData ? console.log("test123") : {}
                            if (ingredientData) {
                                const a = ingredientData.find(x => x.id == mat.ingredient_id)
                                if (a) {
                                    mat.name = a.name;
                                    mat.unit2 = a.unit;
                                    mat.unit2_mass_in_grams = a.mass_in_grams;
                                    mat.substitutes = a.substitutes;
                                }
                                if (unitData) {
                                    const upc1 = unitData.get(mat.unit)
                                    const upc2 = unitData.get(mat.unit2)
                                    if (upc1 && upc2) {
                                        mass = mat.unit2_mass_in_grams * upc2 * mat.amount / upc1
                                    }
                                }
                            }

                            // const mass = mat.unit2_mass_in_grams * 

                            const quality = (mat.quality == "") ? "" : `, ${mat.quality}`;
                            const note = (mat.note == "") ? "" : `; ${mat.note}`;
                            if (mass > 0) {
                                return(
                                    <div key={mat.ingredient_id} className="RecipePageItem IngredientItem">{mat.amount} {mat.unit} {mat.name}{quality}{note} ({mass}g)</div>
                                )
                            } else {
                                return(
                                    <div key={mat.ingredient_id} className="RecipePageItem IngredientItem">{mat.amount} {mat.unit} {mat.name} ( - g)</div>
                                )
                            }
                        })}    
                    </div>
                    <br />
                    <div className="RecipePageBox DirectionsBox">
                        <div className="RecipePageTitle DirectionsTitle">Directions</div>            
                        {recipeData.steps.map(step => {
                            // console.log(step);
                            return(
                                <div key={step.step_num} className="RecipePageItem DirectionsItem">{step.step_num}. {step.text}</div>
                            )
                        })}
                    </div>
                </div> 
            : <div>Loading...</div>
        }
        
    </div> 
    );
}

interface Recipe {
    id: string;
    title: string;
    created: string;
    active_time: number;
    serves_high: number;
    serves_low: number;
    total_time: number;
    author: string;
    materials: Material[]
    steps: Step[]
}

interface Material {
    amount: number;
    ingredient_id: string;
    material_num: number;
    note: string;
    optional: boolean;
    quality: string;
    unit: string;
    name: string;
    unit2: string;
    unit2_mass_in_grams: number;
    substitutes: string[];
}

interface Step {
    ingredient_ids: string[];
    step_num: number;
    text: string;
}

interface Ingredient {
    id: string;
    name: string;
    unit: string;
    mass_in_grams: number;
    substitutes: string[];

}

// interface UnitDict {
    // Dict: Map<string, number>;
    // [key: string]: number;
    // dict: any;
// }

// interface Unit {
//     [key: string]: number;
// }

export default RecipePage;