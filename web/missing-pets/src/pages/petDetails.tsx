import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import "./petDetails.css";

class PetData {
    public name: string;
    public size: string;
    public color: string;
    public breed: string;
    public lastSeen: string;
    public apiHash: string;
};

class OwnerInfo {
    public name: string;
    public phone: string;
}

const PetDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>(); 
  const navigate = useNavigate(); 

  const [pet, setPetInfo] = useState(new PetData());
  const [ownerInfo, setOwnerInfo] = useState(new OwnerInfo());

  useEffect(() => {
      fetch(`/api/pets/${id}`, {method: "GET"}).then(data => data.json<PetData>()).then(final => setPetInfo(final.data));
  }, [id]);

  useEffect(() => {
      fetch(`/api/users/${id}`, {method: "GET"}).then(data => data.json<PetData>()).then(final => setOwnerInfo(final.data));
  }, [id]);

  const handleContactClick = () => {
    alert(`Contactando al dueño: ${ownerInfo.phone}`);
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
        <img
          src={`/images/${pet.apiHash}`}
          alt={`Imagen de ${pet.name}`}
          className="pet-image"
        />
      </div>
      <div className="pet-info">
        <p>
          <strong>Último Lugar Visto:</strong> {pet.lastSeen}
        </p>
        <p>
          <strong>Raza:</strong> {pet.type}
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
