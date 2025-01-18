import React, { useState } from "react";
import "./SearchBar.css";

interface SearchBarProps {
  onSearch: (searchTerm: string) => void; // Función que se ejecutará al buscar
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch }) => {
  const [searchTerm, setSearchTerm] = useState("");

  const handleSearch = () => {
    onSearch(searchTerm); // Llama a la función pasada como prop con el término de búsqueda
  };

  return (
    <div className="search-bar">
      <input
        type="text"
        placeholder="Buscar por ubicación"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      />
      <button onClick={handleSearch}>Filtrar</button>
    </div>
  );
};

export default SearchBar;
