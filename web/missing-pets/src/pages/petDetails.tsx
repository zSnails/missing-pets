import React from "react";
import { useParams, useNavigate } from "react-router-dom";
import "./petDetails.css";

const mockPets = [
  {
    id: "1",
    name: "Luna",
    lastSeenLocation: "Parque Central, Ciudad de México",
    breed: "Labrador Retriever",
    color: "Amarillo",
    size: "Grande",
    images: [
      "https://images.unsplash.com/photo-1574158622682-e40e69881006",
      "https://images.unsplash.com/photo-1517423440428-a5a00ad493e8",
    ],
    ownerContact: "555-123-4567",
  },
  {
    id: "2",
    name: "Max",
    lastSeenLocation: "Calle Los Olivos, Buenos Aires",
    breed: "Pastor Alemán",
    color: "Negro y Marrón",
    size: "Grande",
    images: [
      "https://images.unsplash.com/photo-1558788353-f76d92427f16",
      "https://images.unsplash.com/photo-1574169208507-8437617483c1",
    ],
    ownerContact: "555-987-6543",
  },
];

const PetDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>(); 
  const navigate = useNavigate(); 
  const pet = mockPets.find((pet) => pet.id === id); 

  if (!pet) {
    return <p>Mascota no encontrada.</p>;
  }

  const handleContactClick = () => {
    alert(`Contactando al dueño: ${pet.ownerContact}`);
  };

  const handleBackClick = () => {
    navigate("/lost-pets"); 
  };

  return (
    <div className="pet-detail-container">
      <button className="back-button" onClick={handleBackClick}>
        ← Volver a Mascotas Perdidas
      </button>

      <h1>{pet.name}</h1>
      <div className="pet-detail-images">
        {pet.images.map((image, index) => (
          <img
            key={index}
            src={image}
            alt={`Imagen ${index + 1} de ${pet.name}`}
            className="pet-image"
          />
        ))}
      </div>
      <div className="pet-info">
        <p>
          <strong>Último Lugar Visto:</strong> {pet.lastSeenLocation}
        </p>
        <p>
          <strong>Raza:</strong> {pet.breed}
        </p>
        <p>
          <strong>Color:</strong> {pet.color}
        </p>
        <p>
          <strong>Tamaño:</strong> {pet.size}
        </p>
      </div>
      <button className="contact-button" onClick={handleContactClick}>
        Contactar al Dueño
      </button>
    </div>
  );
};

export default PetDetail;
