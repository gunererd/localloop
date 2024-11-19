CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    parent_id UUID REFERENCES categories(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE category_fields (
    category_id UUID REFERENCES categories(id),
    field_id UUID REFERENCES fields(id),
    is_required BOOLEAN DEFAULT false,
    display_order INTEGER NOT NULL,
    PRIMARY KEY (category_id, field_id)
); 