import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import ToDoList from './ToDoList'; // Adjust the path if necessary

function App() {
  return (
      <Router basename="/todo">
        <Routes>
          <Route path="/" element={<ToDoList />} />
        </Routes>
      </Router>
  );
}

export default App;
