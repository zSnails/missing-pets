import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./components/Layout/Navbar"; // Navbar
import Home from "./pages/home";
import LostPets from "./pages/lostPets";
import Login from "./components/Auth/login";
import Register from "./components/Auth/register";
import AddPet from "./pages/addPet";
function App() {
  return (
    <Router>
      <Navbar />
      <div className="container">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/lost-pets" element={<LostPets />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/add-pet" element={<AddPet />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
