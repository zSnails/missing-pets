import React from "react";
import { useNavigate } from "react-router-dom";
import "./Card.css";

interface CardProps {
  id: string;
  image: string;
  title: string;
  description: string;
  details: string[];
}

const Card: React.FC<CardProps> = ({ id, image, title, description, details }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(`/pet-detail/${id}`);
  };

  return (
    <div className="card" onClick={handleClick}>
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
