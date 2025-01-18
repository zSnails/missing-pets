import React, { useState } from "react";
import "./LostPets.css";
import SearchBar from "../components/Shared/searchBar";
import Card from "../components/Shared/Card"; 

// Datos de prueba
const mockPets = [
  {
    id: "1",
    name: "Luna",
    lastSeenLocation: "Parque Central, Ciudad de México",
    breed: "Labrador Retriever",
    color: "Amarillo",
    size: "Grande",
    image: "https://images.unsplash.com/photo-1574158622682-e40e69881006"
  },
  {
    id: "2",
    name: "Max",
    lastSeenLocation: "Calle Los Olivos, Buenos Aires",
    breed: "Pastor Alemán",
    color: "Negro y Marrón",
    size: "Grande",
    image: "https://images.unsplash.com/photo-1517423440428-a5a00ad493e8"
  },
  {
    id: "3",
    name: "Bella",
    lastSeenLocation: "Plaza Independencia, Montevideo",
    breed: "Beagle",
    color: "Blanco y Marrón",
    size: "Mediano",
    image: "https://images.unsplash.com/photo-1558788353-f76d92427f16"
  }
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
      {/* Usamos el componente de búsqueda */}
      <SearchBar onSearch={handleSearch} />
      <div className="pet-list">
        {filteredPets.length > 0 ? (
          filteredPets.map((pet) => (
            <Card
              key={pet.id}
              image={pet.image}
              title={pet.name}
              description={`Última ubicación: ${pet.lastSeenLocation}`}
              details={[
                `Raza: ${pet.breed}`,
                `Color: ${pet.color}`,
                `Tamaño: ${pet.size}`
              ]}
            />
          ))
        ) : (
          <p>No se encontraron mascotas.</p>
        )}
      </div>
    </div>
  );
};

export default LostPets;
