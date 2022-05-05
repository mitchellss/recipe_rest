import React from 'react';
import logo from './logo.svg';
import './styles/App.css';
import RecipeList from './components/RecipeList';
import {Route, Routes} from "react-router-dom";
import RecipeHome from './components/RecipeHome';
import RecipePage from './components/RecipePage';
import AddRecipePage from './components/AddRecipePage';

function App() {
  return (
      <Routes>
        <Route path="/" element={<RecipeHome/>}/>
        <Route path="/recipe/:recipeId" element={<RecipePage/>}/>
        <Route path="/recipe/add" element={<AddRecipePage/>}/>
      </Routes>
  );
}

export default App;
