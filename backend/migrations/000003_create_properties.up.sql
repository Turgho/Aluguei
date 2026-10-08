CREATE TABLE properties (
    id                    UUID         PRIMARY KEY DEFAULT uuidv7(),
    owner_id              UUID         NOT NULL,
    address_id            UUID         NOT NULL,
    title                 VARCHAR(120) NOT NULL,
    description           TEXT,
    type                  VARCHAR(50)  NOT NULL,
    status                VARCHAR(50)  NOT NULL DEFAULT 'available',
    price_cents           BIGINT       NOT NULL,
    condominium_fee_cents BIGINT       NOT NULL DEFAULT 0,
    iptu_cents            BIGINT       NOT NULL DEFAULT 0,
    bedrooms              INTEGER      NOT NULL DEFAULT 0,
    bathrooms             INTEGER      NOT NULL DEFAULT 0,
    suites                INTEGER      NOT NULL DEFAULT 0,
    parking_spaces        INTEGER      NOT NULL DEFAULT 0,
    area_m2               INTEGER      NOT NULL DEFAULT 0,
    furnished             BOOLEAN      NOT NULL DEFAULT FALSE,
    pet_friendly          BOOLEAN      NOT NULL DEFAULT FALSE,
    has_balcony           BOOLEAN      NOT NULL DEFAULT FALSE,
    has_elevator          BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at            TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at            TIMESTAMPTZ,

    CONSTRAINT fk_properties_owner
        FOREIGN KEY (owner_id) REFERENCES users(id),

    CONSTRAINT fk_properties_address
        FOREIGN KEY (address_id) REFERENCES addresses(id)
);

CREATE INDEX idx_properties_owner_id ON properties(owner_id);
CREATE INDEX idx_properties_address_id ON properties(address_id);
CREATE INDEX idx_properties_type ON properties(type);
CREATE INDEX idx_properties_status ON properties(status);
CREATE INDEX idx_properties_price_cents ON properties(price_cents);
CREATE INDEX idx_properties_deleted_at ON properties(deleted_at);