CREATE TABLE field_types (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type_discriminator_id UUID NOT NULL REFERENCES field_type_discriminators(id),
    properties JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE fields (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    field_type_id UUID NOT NULL REFERENCES field_types(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
); 