import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./Navbar.css";

const Navbar: React.FC = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen); 
  };

  return (
    <nav className="navbar">
      <div className="navbar-logo">
        <Link to="/">🐾 Mascotas Perdidas</Link>
      </div>

      <button className="navbar-toggle" onClick={toggleMenu}>
        ☰
      </button>

      <ul className={`navbar-links ${isMenuOpen ? "navbar-links-active" : ""}`}>
        <li>
          <Link to="/lost-pets" onClick={() => setIsMenuOpen(false)}>
            Ver Mascotas Perdidas
          </Link>
        </li>
        <li>
          <Link to="/Add-pet" onClick={() => setIsMenuOpen(false)}>
            Publicar Mascota
          </Link>
        </li>
        <li>
          <Link to="/login" onClick={() => setIsMenuOpen(false)}>
            Inicio de sesión
          </Link>
        </li>
      </ul>
    </nav>
  );
};

export default Navbar;
