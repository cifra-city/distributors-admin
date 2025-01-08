CREATE TABLE distributors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    owner_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE distributors_employees (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    distributors_id UUID NOT NULL REFERENCES distributors(id) ON DELETE CASCADE,
    user_id UUID NOT NULL,
    role roles NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TYPE roles AS ENUM ('owner', 'admin', 'moderator', 'staff', 'member');

CREATE UNIQUE INDEX unique_owner_per_distributor
    ON distributors_employees (distributors_id)
    WHERE role = 'owner';
