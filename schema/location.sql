CREATE TABLE building (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR UNIQUE NOT NULL CHECK (1 < char_length(name) AND char_length(name) < 256)
);

CREATE TABLE location (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    building_id UUID NOT NULL,
    name VARCHAR NOT NULL CHECK (1 < char_length(name) AND char_length(name) < 256),
    UNIQUE(building_id, name),
    FOREIGN KEY (building_id) REFERENCES building(id)
);
