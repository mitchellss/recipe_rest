import axios from "axios";
import { useEffect, useState } from "react";

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
                return(
                    <div key={test.id}>{test.title}</div>
                )
            }) : <div>No data yet</div>
        }
    </div>
    );
}

export default RecipeList;