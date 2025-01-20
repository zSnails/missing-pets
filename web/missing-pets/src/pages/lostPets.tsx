import React, { useState, useEffect } from "react";
import "./lostPets.css";
import SearchBar from "../components/Shared/SearchBar";
import Card from "../components/Shared/Card";

const mockPets = [
  {
    id: "1",
    name: "Luna",
    lastSeen: "Parque Central, Ciudad de México",
    type: "Labrador Retriever",
    color: "Amarillo",
    size: "Grande",
    image: ["https://images.unsplash.com/photo-1574158622682-e40e69881006"],
  },
  {
    id: "2",
    name: "Max",
    lastSeen: "Calle Los Olivos, Buenos Aires",
    type: "Pastor Alemán",
    color: "Negro y Marrón",
    size: "Grande",
    image: ["https://images.unsplash.com/photo-1558788353-f76d92427f16"],
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

  useEffect(() => {
      fetch("/api/pets", { method: "GET"}).then(data => data.json()).then(final => { setFilteredPets(final.data); console.log(final)});
  }, []);

  return (
    <div className="lost-pets-container">
      <h1>Mascotas Perdidas</h1>
      <SearchBar onSearch={handleSearch} />
      <div className="pet-list">
        {filteredPets.map((pet) => (
          <Card
            key={pet.id}
            id={pet.id}
            image={`/images/${pet.apiHash}`}
            title={pet.name}
            description={`Último lugar visto: ${pet.lastSeen}`}
            details={[`Raza: ${pet.type}`, `Color: ${pet.color}`, `Tamaño: ${pet.size}`]}
          />
        ))}
      </div>
    </div>
  );
};

export default LostPets;
