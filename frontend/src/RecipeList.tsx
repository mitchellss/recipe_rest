import axios from "axios";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import RecipePage from "./RecipePage";

function RecipeList() {
    const [data, setData] = useState<any[]>([]);

    const getAllData = () => {
        axios.get("http://192.168.1.125:8080/api/recipe")
            .then(res => {
                setData(res.data);
                console.log(res.data)
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
        {data ?
            data.map(test => {
                const recipePage = <RecipePage key={test.id} recipe={test}></RecipePage>
                return(
                    <div>
                        <Link to={`/${test.id}`}>
                        </Link>
                    </div>
                )
            }) : <div>No data yet</div>
        }
    </div>
    );
}

export default RecipeList;