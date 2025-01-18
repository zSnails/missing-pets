import React, { useState } from "react";
import "./LostPets.css";
import SearchBar from "../components/Shared/SearchBar";
import Card from "../components/Shared/Card";

const mockPets = [
  {
    id: "1",
    name: "Luna",
    lastSeenLocation: "Parque Central, Ciudad de México",
    breed: "Labrador Retriever",
    color: "Amarillo",
    size: "Grande",
    image: "https://images.unsplash.com/photo-1574158622682-e40e69881006",
  },
  {
    id: "2",
    name: "Max",
    lastSeenLocation: "Calle Los Olivos, Buenos Aires",
    breed: "Pastor Alemán",
    color: "Negro y Marrón",
    size: "Grande",
    image: "https://images.unsplash.com/photo-1558788353-f76d92427f16",
  },
];

const LostPets: React.FC = () => {
  const [filteredPets, setFilteredPets] = useState(mockPets);

  const handleSearch = (searchTerm: string) => {
    const filtered = mockPets.filter((pet) =>
      pet.lastSeenLocation.toLowerCase().includes(searchTerm.toLowerCase())
    );
    setFilteredPets(filtered);
  };

  return (
    <div className="lost-pets-container">
      <h1>Mascotas Perdidas</h1>
      <SearchBar onSearch={handleSearch} />
      <div className="pet-list">
        {filteredPets.map((pet) => (
          <Card
            key={pet.id}
            id={pet.id}
            image={pet.image}
            title={pet.name}
            description={`Último lugar visto: ${pet.lastSeenLocation}`}
            details={[`Raza: ${pet.breed}`, `Color: ${pet.color}`, `Tamaño: ${pet.size}`]}
          />
        ))}
      </div>
    </div>
  );
};

export default LostPets;
