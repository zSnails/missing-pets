import React, { useState, useEffect } from "react";
import "../Pets/PetList.css";

// Datos de prueba en formato JSON
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
  },
  {
    id: "4",
    name: "Rocky",
    lastSeenLocation: "Parque del Retiro, Madrid",
    breed: "Bulldog Inglés",
    color: "Blanco y Negro",
    size: "Mediano",
    image: "https://images.unsplash.com/photo-1583511655626-9b7eaf6e1216"
  },
  {
    id: "5",
    name: "Simba",
    lastSeenLocation: "Avenida Paulista, São Paulo",
    breed: "Golden Retriever",
    color: "Dorado",
    size: "Grande",
    image: "https://images.unsplash.com/photo-1525253086316-d0c936c814f8"
  }
];

const PetList: React.FC = () => {
  const [pets, setPets] = useState(mockPets); // Usamos los datos de prueba como estado inicial
  const [filteredPets, setFilteredPets] = useState(mockPets);
  const [locationFilter, setLocationFilter] = useState("");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");

  // Función para filtrar las mascotas según los filtros activos
  const handleFilter = () => {
    let filtered = pets;

    if (locationFilter) {
      filtered = filtered.filter((pet) =>
        pet.lastSeenLocation.toLowerCase().includes(locationFilter.toLowerCase())
      );
    }

    setFilteredPets(filtered);
  };

  return (
    <div className="container">
      <h2 className="title">Mascotas Perdidas</h2>

      {/* Barra de búsqueda */}
      <div className="filters">
        <div className="filter-group">
          <label>Buscar por lugar</label>
          <input
            type="text"
            value={locationFilter}
            onChange={(e) => setLocationFilter(e.target.value)}
            placeholder="Ej. Ciudad, Barrio"
          />
        </div>
        <button onClick={handleFilter} className="filter-button">
          Filtrar
        </button>
      </div>

      {/* Listado de mascotas */}
      <div className="pet-list">
        {filteredPets.length > 0 ? (
          filteredPets.map((pet) => (
            <div key={pet.id} className="pet-card">
              <img src={pet.image} alt={pet.name} className="pet-image" />
              <h3>{pet.name}</h3>
              <p><strong>Raza:</strong> {pet.breed}</p>
              <p><strong>Color:</strong> {pet.color}</p>
              <p><strong>Tamaño:</strong> {pet.size}</p>
              <p className="pet-meta">
                <strong>Última ubicación:</strong> {pet.lastSeenLocation}
              </p>
            </div>
          ))
        ) : (
          <p>No se encontraron resultados.</p>
        )}
      </div>
    </div>
  );
};

export default PetList;
