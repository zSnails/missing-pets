import React from "react";
import { Link } from "react-router-dom";
import "./home.css";

const Home: React.FC = () => {
  return (
    <div className="home-wrapper">
      <div className="home-container">
        <header className="home-header">
          <h1>¡Ayuda a encontrar mascotas perdidas!</h1>
          <p>
            Conecta con dueños y encuentra a los amigos peludos que necesitan regresar a casa.
          </p>
          <div className="home-buttons">
            <Link to="/lost_pets" className="home-button">
              Ver Mascotas Perdidas
            </Link>
            <Link to="/register" className="home-button">
              Registrate
            </Link>
          </div>
        </header>
        <section className="home-section">
          <h2>¿Cómo funciona?</h2>
          <div className="home-steps">
            <div className="step">
              <h3>🔍 Busca</h3>
              <p>Explora publicaciones de mascotas perdidas en tu área.</p>
            </div>
            <div className="step">
              <h3>📢 Publica</h3>
              <p>Crea una publicación si tu mascota está perdida.</p>
            </div>
            <div className="step">
              <h3>🤝 Conecta</h3>
              <p>Ponte en contacto con dueños para ayudar a encontrar mascotas.</p>
            </div>
          </div>
        </section>
      </div>
    </div>
  );
};

export default Home;
