CREATE TABLE IF NOT EXISTS contact_us (
    id SERIAL PRIMARY KEY,
    company_name VARCHAR(150),
    location_name VARCHAR(150),
    address text,
    phone_number VARCHAR(17),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP

);