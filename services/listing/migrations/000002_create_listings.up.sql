CREATE TABLE listings (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category_id UUID NOT NULL,
    price DECIMAL,
    currency_id UUID REFERENCES currencies(id),
    condition_id UUID REFERENCES conditions(id),
    status_id UUID NOT NULL REFERENCES listing_statuses(id),
    media_url VARCHAR(500),
    custom_fields JSONB NOT NULL DEFAULT '{}',
    created_by UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    published_at TIMESTAMP
);

CREATE INDEX idx_listings_category ON listings(category_id);
CREATE INDEX idx_listings_status ON listings(status_id);
CREATE INDEX idx_listings_created_by ON listings(created_by);
CREATE INDEX idx_listings_price ON listings(price);
CREATE INDEX idx_listing_custom_fields ON listings USING GIN(custom_fields); 