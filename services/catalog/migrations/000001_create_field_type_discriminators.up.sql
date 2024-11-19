CREATE TABLE field_type_discriminators (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    validation_schema JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
); 