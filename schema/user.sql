CREATE TABLE "user" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    local_id VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR NOT NULL CHECK (1 < char_length(display_name) AND char_length(display_name) < 256),
    username VARCHAR NOT NULL CHECK (1 < char_length(username) AND char_length(username) < 256)
);

CREATE TABLE location_user_assignment (
    location_id UUID NOT NULL,
    user_id UUID NOT NULL,
    PRIMARY KEY(location_id, user_id),
    FOREIGN KEY (location_id) REFERENCES location(id),
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

CREATE TABLE "group" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    local_id VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR NOT NULL CHECK (1 < char_length(display_name) AND char_length(display_name) < 256)
);

CREATE TABLE group_user_assignment (
    "group_id" UUID NOT NULL,
    user_id UUID NOT NULL,
    PRIMARY KEY("group_id", user_id),
    FOREIGN KEY ("group_id") REFERENCES "group"(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

CREATE TABLE location_group_assignment (
    location_id UUID NOT NULL,
    "group_id" UUID NOT NULL,
    PRIMARY KEY(location_id, "group_id"),
    FOREIGN KEY (location_id) REFERENCES location(id),
    FOREIGN KEY ("group_id") REFERENCES "group"(id) ON DELETE CASCADE
);

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
