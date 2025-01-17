CREATE TABLE pet_owners (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    phone TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    address TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS pet_owners_email_idx ON pet_owners (email);

CREATE TABLE missing_pet_owner_rel (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    missing_pet_id INT NOT NULL,
    pet_owner_id INT NOT NULL,
    FOREIGN KEY (missing_pet_id) REFERENCES missing_pets(id) ON DELETE CASCADE ON UPDATE CASCADE
    FOREIGN KEY (pet_owner_id) REFERENCES pet_owners(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX IF NOT EXISTS missing_pet_owner_idx ON missing_pet_owner_rel (missing_pet_id, pet_owner_id);

CREATE TABLE missing_pets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    last_seen TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS missing_pets_idx ON missing_pets (name, type, last_seen);

CREATE TABLE missing_pet_photos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    pet_id INT NOT NULL,
    encoded_data TEXT NOT NULL,
    FOREIGN KEY (pet_id) REFERENCES missing_pets(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- TODO: create posts table, cuz there have to be posts, this will work more as
-- a forum than anything else
