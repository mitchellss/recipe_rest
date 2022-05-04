import axios from "axios";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { HOST, PORT } from "./constants";

function RecipeList() {
    const [data, setData] = useState<any[]>([]);

    const getAllData = () => {
        axios.get(`${HOST}:${PORT}/api/recipe`)
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
            <h1 className="Title Large">Recipes:</h1>
            {data ?
                data.map(test => {
                    return (
                        <Link key={test.id} to={`/recipe/${test.id}`}>
                            <div className="RecipeListing">{test.title}</div>
                        </Link>
                    )
                }) : <div>No data yet</div>
            }
        </div>
    );
}

export default RecipeList;