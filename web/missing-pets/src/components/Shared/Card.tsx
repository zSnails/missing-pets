import React from "react";
import "./Card.css";

interface CardProps {
  image: string; // URL de la imagen
  title: string; // Título principal
  description: string; // Descripción breve
  details: string[]; // Detalles adicionales (lista de textos)
}

const Card: React.FC<CardProps> = ({ image, title, description, details }) => {
  return (
    <div className="card">
      <img src={image} alt={title} className="card-image" />
      <div className="card-content">
        <h2 className="card-title">{title}</h2>
        <p className="card-description">{description}</p>
        <ul className="card-details">
          {details.map((detail, index) => (
            <li key={index}>{detail}</li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default Card;
