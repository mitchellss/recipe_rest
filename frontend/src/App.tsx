import React from 'react';
import logo from './logo.svg';
import './App.css';
import RecipeList from './RecipeList';
import {Route, Routes} from "react-router-dom";
import RecipeHome from './RecipeHome';
import RecipePage from './RecipePage';

function App() {
  return (
      <Routes>
        <Route path="/" element={<RecipeHome/>}/>
        <Route path="/recipe/:recipeId" element={<RecipePage/>}/>
      </Routes>
  );
}

export default App;
