CREATE TABLE IF NOT EXISTS hero_sections (
    id SERIAL PRIMARY KEY,
    heading VARCHAR(150),
    sub_heading VARCHAR(150),
    path_video text NULL,
    path_banner text NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP

);