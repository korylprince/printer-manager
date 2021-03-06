CREATE TABLE manufacturer (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR UNIQUE NOT NULL CHECK (1 <= char_length(name) AND char_length(name) < 256)
);

CREATE TABLE model (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    manufacturer_id UUID NOT NULL,
    name VARCHAR NOT NULL CHECK (1 <= char_length(name) AND char_length(name) < 256),
    driver JSONB NOT NULL,
    UNIQUE(manufacturer_id, name),
    FOREIGN KEY (manufacturer_id) REFERENCES manufacturer(id)
);

CREATE TABLE printer (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    model_id UUID NOT NULL,
    location_id UUID NOT NULL,
    hostname VARCHAR NOT NULL CHECK (1 <= char_length(hostname) AND char_length(hostname) < 256),
    driver_extra JSONB,
    UNIQUE(model_id, location_id, hostname),
    FOREIGN KEY (model_id) REFERENCES model(id),
    FOREIGN KEY (location_id) REFERENCES location(id)
);

CREATE VIEW printer_flattened AS
    SELECT
        printer.id AS id,
        printer.hostname AS hostname,
        format('%s %s %s', location.name, manufacturer.name, model.name) AS name,
        format('%s - %s', building.name, location.name) AS location,
        COALESCE(jsonb_merge(model.driver, printer.driver_extra), '{}'::jsonb) AS driver
    FROM printer
    INNER JOIN model ON
        printer.model_id = model.id
    INNER JOIN manufacturer ON
        model.manufacturer_id = manufacturer.id
    INNER JOIN location ON
        printer.location_id = location.id
    INNER JOIN building ON
        location.building_id = building.id;

CREATE VIEW user_printers AS
    SELECT
        assigned.user_id,
        printer_flattened.*
    FROM printer_flattened INNER JOIN ((
        SELECT
            "user".id AS user_id,
            printer.id AS printer_id
        FROM "user"
        INNER JOIN location_user_assignment ON
            "user".id = location_user_assignment.user_id
        INNER JOIN location ON
            location_user_assignment.location_id = location.id
        INNER JOIN printer ON
            location.id = printer.location_id
    ) UNION (
        SELECT
            "user".id AS user_id,
            printer.id AS printer_id
        FROM "user"
        INNER JOIN group_user_assignment ON
            "user".id = group_user_assignment.user_id
        INNER JOIN "group" ON
            group_user_assignment."group_id" = "group".id
        INNER JOIN location_group_assignment ON
            "group".id = location_group_assignment."group_id"
        INNER JOIN location ON
            location_group_assignment.location_id = location.id
        INNER JOIN printer ON
            location.id = printer.location_id
    )) AS assigned ON
        printer_flattened.id = assigned.printer_id;
