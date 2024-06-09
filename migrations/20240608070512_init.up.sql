CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS coffees (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT(uuid_generate_v4 ()),
    "name" varchar NOT NULL,
    "roast" varchar NOT NULL,
    "region" varchar NOT NULL,
    "image" varchar NOT NULL,
    "price" Float NOT NUll,
    "grind_unit" INT NOT NUll,
    "create_at" TIMESTAMP
    WITH
        TIME ZONE DEFAULT NOW(),
        "update_at" TIMESTAMP
    WITH
        TIME ZONE DEFAULT NOW()
);