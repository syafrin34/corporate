CREATE TABLE IF NOT EXISTS portofolio_details (
    id SERIAL PRIMARY KEY,
    portofolio_sections_id INT REFERENCES portofolio_sections(id) ON DELETE CASCADE,
    category VARCHAR(150),
    client_name text,
    project_date TIMESTAMP,
    project_url text NULL,
    title VARCHAR(200) ,
    description text,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP

);
CREATE INDEX idx_portofolio_details_portofolio_section_id ON portofolio_details(portofolio_sections_id);