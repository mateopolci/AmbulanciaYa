---------------------------------------------------------------------------------------------------------------------------------

-- TABLES

-- Create extension for UUID support
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create table choferes
CREATE TABLE choferes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nombreCompleto VARCHAR(255) NOT NULL,
    dni VARCHAR(20) NOT NULL
);

-- Create table paramedicos
CREATE TABLE paramedicos (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nombreCompleto VARCHAR(255) NOT NULL,
    dni VARCHAR(60) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(20) NOT NULL,
    isAdmin BOOLEAN NOT NULL
);

-- Create table ambulancias
CREATE TABLE ambulancias (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patente VARCHAR(20) NOT NULL,
    inventario BOOLEAN NOT NULL,
    vtv BOOLEAN NOT NULL,
    seguro BOOLEAN NOT NULL,
    choferID UUID NOT NULL,
    paramedicoID UUID NOT NULL,
    FOREIGN KEY (choferID) REFERENCES choferes(id),
    FOREIGN KEY (paramedicoID) REFERENCES paramedicos(id)
    base BOOLEAN NOT NULL,
    cadenas BOOLEAN NOT NULL,
    antinieblas BOOLEAN NOT NULL,
    cubiertasLluvia BOOLEAN NOT NULL,
);

-- Create table hospitales
CREATE TABLE hospitales (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nombre VARCHAR(255) NOT NULL,
    direccion VARCHAR(255) NOT NULL
);

-- Create table pacientes
CREATE TABLE pacientes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nombreCompleto VARCHAR(255) NOT NULL,
    telefono VARCHAR(20) NOT NULL
);

-- Create table accidentes
CREATE TABLE accidentes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    direccion VARCHAR(255) NOT NULL,
    descripcion TEXT NOT NULL,
    fecha VARCHAR(10) NOT NULL,
    hora VARCHAR(8) NOT NULL,
    ambulanciaID UUID NOT NULL,
    hospitalID UUID,
    pacienteID UUID,
    FOREIGN KEY (ambulanciaID) REFERENCES ambulancias(id),
    FOREIGN KEY (hospitalID) REFERENCES hospitales(id),
    FOREIGN KEY (pacienteID) REFERENCES pacientes(id)
);

-- Create table reportes
CREATE TABLE reportes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    descripcion TEXT NOT NULL,
    fecha VARCHAR(10) NOT NULL,
    hora VARCHAR(8) NOT NULL,
    requiereTraslado BOOLEAN NOT NULL,
    accidenteId UUID NOT NULL,
    FOREIGN KEY (accidenteId) REFERENCES accidentes(id)
);

---------------------------------------------------------------------------------------------------------------------------------

-- CONSTRAINTS

-- Constraint para Accidente-Reporte (1:0..1)
-- Asegura que un accidente no puede tener más de un reporte
ALTER TABLE reportes
ADD CONSTRAINT unique_accidente_reporte
UNIQUE (accidenteId);

-- Añadir constraint de unicidad para chofer-ambulancia (1:1)
ALTER TABLE ambulancias
ADD CONSTRAINT unique_chofer_ambulancia 
UNIQUE (choferID);

-- Añadir constraint de unicidad para paramedico-ambulancia (1:1)
ALTER TABLE ambulancias
ADD CONSTRAINT unique_paramedico_ambulancia 
UNIQUE (paramedicoID);

-- Añadir constraint de unicidad para accidente-paciente (1..*:0..1) 
ALTER TABLE accidentes
ADD CONSTRAINT accidentes_pacienteid_fkey 
FOREIGN KEY (pacienteID) 
REFERENCES pacientes(id)
ON DELETE SET NULL;

---------------------------------------------------------------------------------------------------------------------------------

-- Test values

-- Datos de prueba para choferes
INSERT INTO choferes (id, nombreCompleto, dni) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'Juan Pérez', '12345678'),
('550e8400-e29b-41d4-a716-446655440001', 'María López', '23456789'),
('550e8400-e29b-41d4-a716-446655440002', 'Carlos Gómez', '34567890'),
('550e8400-e29b-41d4-a716-446655440003', 'Ana Torres', '45678901'),
('550e8400-e29b-41d4-a716-446655440004', 'Luis Ramírez', '56789012');

-- Datos de prueba para paramedicos
INSERT INTO paramedicos (id, nombreCompleto, dni) VALUES
('550e8400-e29b-41d4-a716-446655440005', 'Pedro Sánchez', '67890123'),
('550e8400-e29b-41d4-a716-446655440006', 'Marta Castillo', '78901234'),
('550e8400-e29b-41d4-a716-446655440007', 'Sofía Ruiz', '89012345'),
('550e8400-e29b-41d4-a716-446655440008', 'Jorge Vargas', '90123456'),
('550e8400-e29b-41d4-a716-446655440009', 'Laura Díaz', '01234567');

-- Datos de prueba para ambulancias
INSERT INTO ambulancias (id, patente, inventario, vtv, seguro, choferID, paramedicoID, base) VALUES
('550e8400-e29b-41d4-a716-446655440010', 'ABC123', TRUE, TRUE, TRUE, '550e8400-e29b-41d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440005', TRUE),
('550e8400-e29b-41d4-a716-446655440011', 'DEF456', TRUE, TRUE, TRUE, '550e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440006', FALSE),
('550e8400-e29b-41d4-a716-446655440012', 'GHI789', FALSE, TRUE, FALSE, '550e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440007', TRUE),
('550e8400-e29b-41d4-a716-446655440013', 'JKL012', TRUE, FALSE, TRUE, '550e8400-e29b-41d4-a716-446655440003', '550e8400-e29b-41d4-a716-446655440008', FALSE),
('550e8400-e29b-41d4-a716-446655440014', 'MNO345', TRUE, TRUE, FALSE, '550e8400-e29b-41d4-a716-446655440004', '550e8400-e29b-41d4-a716-446655440009', FALSE);

-- Datos de prueba para hospitales
INSERT INTO hospitales (id, nombre, direccion) VALUES
('550e8400-e29b-41d4-a716-446655440015', 'Hospital Central', 'Av. Principal 123'),
('550e8400-e29b-41d4-a716-446655440016', 'Clínica San José', 'Calle Secundaria 456'),
('550e8400-e29b-41d4-a716-446655440017', 'Hospital del Norte', 'Av. Libertad 789'),
('550e8400-e29b-41d4-a716-446655440018', 'Clínica del Sur', 'Calle Paz 101'),
('550e8400-e29b-41d4-a716-446655440019', 'Hospital Universitario', 'Av. Sabiduría 202');

-- Datos de prueba para pacientes
INSERT INTO pacientes (id, nombreCompleto, telefono) VALUES
('550e8400-e29b-41d4-a716-446655440020', 'Lucía Fernández', '111111111'),
('550e8400-e29b-41d4-a716-446655440021', 'Ricardo Morales', '222222222'),
('550e8400-e29b-41d4-a716-446655440022', 'Gabriela Torres', '333333333'),
('550e8400-e29b-41d4-a716-446655440023', 'Diego Castro', '444444444'),
('550e8400-e29b-41d4-a716-446655440024', 'Carla Benítez', '555555555');

-- Datos de prueba para accidentes
INSERT INTO accidentes (id, direccion, descripcion, fecha, hora, ambulanciaID, hospitalID, pacienteID) VALUES
('550e8400-e29b-41d4-a716-446655440025', 'Av. Siempreviva 742', 'Colisión múltiple', '2025-01-01', '08:00', '550e8400-e29b-41d4-a716-446655440010', '550e8400-e29b-41d4-a716-446655440015', '550e8400-e29b-41d4-a716-446655440020'),
('550e8400-e29b-41d4-a716-446655440026', 'Calle Falsa 123', 'Atropello', '2025-01-02', '09:30', '550e8400-e29b-41d4-a716-446655440011', '550e8400-e29b-41d4-a716-446655440016', '550e8400-e29b-41d4-a716-446655440021'),
('550e8400-e29b-41d4-a716-446655440027', 'Av. Colón 456', 'Vuelco de vehículo', '2025-01-03', '10:15', '550e8400-e29b-41d4-a716-446655440012', NULL, '550e8400-e29b-41d4-a716-446655440022'),
('550e8400-e29b-41d4-a716-446655440028', 'Calle Corrientes 789', 'Incendio en vivienda', '2025-01-04', '11:00', '550e8400-e29b-41d4-a716-446655440013', '550e8400-e29b-41d4-a716-446655440017', '550e8400-e29b-41d4-a716-446655440023'),
('550e8400-e29b-41d4-a716-446655440029', 'Ruta Nacional 40', 'Accidente de camión', '2025-01-05', '12:45', '550e8400-e29b-41d4-a716-446655440014', NULL, '550e8400-e29b-41d4-a716-446655440024');

-- Datos de prueba para reportes
INSERT INTO reportes (id, descripcion, fecha, hora, requiereTraslado, accidenteId) VALUES
('550e8400-e29b-41d4-a716-446655440030', 'Paciente estable, traslado al hospital', '2025-01-01', '08:30', TRUE, '550e8400-e29b-41d4-a716-446655440025'),
('550e8400-e29b-41d4-a716-446655440031', 'Paciente con fracturas, traslado urgente', '2025-01-02', '10:00', TRUE, '550e8400-e29b-41d4-a716-446655440026'),
('550e8400-e29b-41d4-a716-446655440032', 'Sin traslado necesario, paciente estable', '2025-01-03', '10:45', FALSE, '550e8400-e29b-41d4-a716-446655440027'),
('550e8400-e29b-41d4-a716-446655440033', 'Paciente con quemaduras, traslado urgente', '2025-01-04', '11:30', TRUE, '550e8400-e29b-41d4-a716-446655440028'),
('550e8400-e29b-41d4-a716-446655440034', 'Paciente fallecido, sin traslado', '2025-01-05', '13:15', FALSE, '550e8400-e29b-41d4-a716-446655440029');


---------------------------------------------------------------------------------------------------------------------------------
-- SELECTS DE TABLAS

-- Accidentes
SELECT *
FROM public.accidentes;

-- Ambulancias
SELECT *
FROM public.ambulancias;

-- Choferes
SELECT *
FROM public.choferes;

-- Hospitales
SELECT *
FROM public.hospitales;

-- Pacientes
SELECT *
FROM public.pacientes;

-- Paramedicos
SELECT *
FROM public.paramedicos;

-- Reportes
SELECT *
FROM public.reportes;

---------------------------------------------------------------------------------------------------------------------------------

-- DANGER ZONE 

-- Drop tables

-- Primero eliminamos las tablas que tienen foreign keys
DROP TABLE IF EXISTS reportes CASCADE;
DROP TABLE IF EXISTS accidentes CASCADE;
DROP TABLE IF EXISTS ambulancias CASCADE;

-- Luego eliminamos las tablas independientes
DROP TABLE IF EXISTS choferes CASCADE;
DROP TABLE IF EXISTS paramedicos CASCADE;
DROP TABLE IF EXISTS hospitales CASCADE;
DROP TABLE IF EXISTS pacientes CASCADE;