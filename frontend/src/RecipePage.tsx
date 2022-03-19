import axios from "axios";
import { useState, useEffect } from "react";
import {useParams} from "react-router-dom";
import internal from "stream";
import { HOST, PORT } from "./constants";


function RecipePage() {
    const params = useParams()
    const [data, setData] = useState<Recipe>();

    const getAllData = () => {
        axios.get(`${HOST}:${PORT}/api/recipe/${params.recipeId}`)
            .then(res => {
                setData(res.data);
                // console.log(res.data)
            })
            .catch((error) => {
                console.log(error);
            });
    }

    useEffect(() => {
        getAllData();
    }, []);

    return ( 
    <div>
        {
            data ? 
                <div>
                    <div>{data.title}</div>
                    <div>{data.author}</div>
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
}

export default RecipePage;