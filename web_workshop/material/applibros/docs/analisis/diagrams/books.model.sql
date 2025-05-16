-- PostgreSQL schema for books database
-- Generated from PlantUML diagram

-- Drop existing tables if they exist
DROP TABLE IF EXISTS autor_genero;
DROP TABLE IF EXISTS imagen;
DROP TABLE IF EXISTS libro;
DROP TABLE IF EXISTS autor;
DROP TABLE IF EXISTS genero;

-- Create tables
CREATE TABLE autor (
    id SERIAL PRIMARY KEY,
    nombre TEXT NOT NULL
);

CREATE TABLE genero (
    id SERIAL PRIMARY KEY,
    nombre TEXT NOT NULL
);

CREATE TABLE libro (
    id TEXT PRIMARY KEY,
    titulo TEXT NOT NULL,
    autor_id INTEGER NOT NULL,
    genero_id INTEGER NOT NULL,
    FOREIGN KEY (autor_id) REFERENCES autor(id),
    FOREIGN KEY (genero_id) REFERENCES genero(id)
);

CREATE TABLE imagen (
    id SERIAL PRIMARY KEY,
    libro_id TEXT NOT NULL,
    url TEXT NOT NULL,
    descripcion TEXT,
    FOREIGN KEY (libro_id) REFERENCES libro(id)
);

-- Junction table for many-to-many relationship between Autor and Genero
CREATE TABLE autor_genero (
    autor_id INTEGER NOT NULL,
    genero_id INTEGER NOT NULL,
    PRIMARY KEY (autor_id, genero_id),
    FOREIGN KEY (autor_id) REFERENCES autor(id),
    FOREIGN KEY (genero_id) REFERENCES genero(id)
);

-- Create indexes for better performance
CREATE INDEX idx_libro_autor ON libro(autor_id);
CREATE INDEX idx_libro_genero ON libro(genero_id);
CREATE INDEX idx_imagen_libro ON imagen(libro_id);
CREATE INDEX idx_autor_genero_autor ON autor_genero(autor_id);
CREATE INDEX idx_autor_genero_genero ON autor_genero(genero_id);
