CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    service_id INT REFERENCES service_sections(id) ON DELETE CASCADE,
    name VARCHAR(150) ,
    phone_number VARCHAR(15),
    email VARCHAR(150),
    brief text,
    meet_at TIMESTAMP NOT NULL,
    budget DECIMAL(10,1) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP

);
CREATE INDEX idx_appointments_service_id ON appointments(service_id);