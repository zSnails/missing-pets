import React, { useState } from "react";
import "./AddPet.css";

const AddPet: React.FC = () => {
  const [petName, setPetName] = useState("");
  const [lastSeenLocation, setLastSeenLocation] = useState("");
  const [breed, setBreed] = useState("");
  const [color, setColor] = useState("");
  const [size, setSize] = useState("");
  const [images, setImages] = useState<File[]>([]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (images.length === 0) {
      alert("Por favor, sube al menos una imagen de la mascota.");
      return;
    }

    const formData = new FormData();
    formData.append("petName", petName);
    formData.append("lastSeenLocation", lastSeenLocation);
    formData.append("breed", breed);
    formData.append("color", color);
    formData.append("size", size);
    images.forEach((image, index) => {
      formData.append(`images[${index}]`, image);
    });

    console.log("Datos de la mascota perdida:", {
      petName,
      lastSeenLocation,
      breed,
      color,
      size,
      images,
    });
  };

  const handleImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const filesArray = Array.from(e.target.files);
      setImages(filesArray);
    }
  };

  return (
    <div className="add-pet-container">
      <h1>Agregar Mascota Perdida</h1>
      <form className="add-pet-form" onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="petName">Nombre de la Mascota</label>
          <input
            type="text"
            id="petName"
            placeholder="Ingresa el nombre"
            value={petName}
            onChange={(e) => setPetName(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="lastSeenLocation">Último Lugar Visto</label>
          <input
            type="text"
            id="lastSeenLocation"
            placeholder="Ingresa el último lugar visto"
            value={lastSeenLocation}
            onChange={(e) => setLastSeenLocation(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="breed">Raza</label>
          <input
            type="text"
            id="breed"
            placeholder="Ingresa la raza"
            value={breed}
            onChange={(e) => setBreed(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="color">Color</label>
          <input
            type="text"
            id="color"
            placeholder="Ingresa el color"
            value={color}
            onChange={(e) => setColor(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="size">Tamaño</label>
          <select
            id="size"
            value={size}
            onChange={(e) => setSize(e.target.value)}
            required
          >
            <option value="">Selecciona un tamaño</option>
            <option value="Pequeño">Pequeño</option>
            <option value="Mediano">Mediano</option>
            <option value="Grande">Grande</option>
          </select>
        </div>
        <div className="form-group file-input-wrapper">
          <label>Imágenes de la Mascota</label>
          <input
            type="file"
            id="images"
            accept="image/*"
            multiple
            onChange={handleImageChange}
          />
          <button type="button" className="custom-file-button">
            Seleccionar Imágenes
          </button>
          <span className="file-selected">
            {images.length > 0
              ? `${images.length} archivo(s) seleccionado(s)`
              : "No se han seleccionado archivos"}
          </span>
        </div>
        <div className="image-preview">
          {images.map((image, index) => (
            <div key={index} className="image-preview-item">
              <img src={URL.createObjectURL(image)} alt={`Preview ${index}`} />
            </div>
          ))}
        </div>
        <button type="submit" className="add-pet-button">
          Agregar Mascota
        </button>
      </form>
    </div>
  );
};

export default AddPet;
