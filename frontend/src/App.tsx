import React from 'react';
import logo from './logo.svg';
import './App.css';
import RecipeList from './RecipeList';
import {Route, Routes} from "react-router-dom";
import RecipeHome from './RecipeHome';
import RecipePage from './RecipePage';

function App() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<RecipeHome/>}/>
        <Route path="/recipe/:recipeId" element={<RecipePage/>}/>
      </Routes>
    </div>
  );
}

export default App;
