import React from 'react';
import logo from './logo.svg';
import './App.css';
import RecipeList from './RecipeList';

function App() {
  return (
    <div className="App">
      <div className='Header'>
        RecipeRest
      </div>
      <RecipeList/>
    </div>
  );
}

export default App;
