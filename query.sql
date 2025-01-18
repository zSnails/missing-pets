-- name: CreateMissingPet :one
INSERT INTO
missing_pets (name, type, last_seen, owner_id)
VALUES (?, ?, ?, ?)
RETURNING id, name, type, last_seen;

-- name: FindMissingPetsByName :many
SELECT id, name, type, last_seen FROM
missing_pets
WHERE remove_special_characters(name)
LIKE remove_special_characters(CAST(sqlc.arg(name) AS TEXT));

-- name: DoesUserOwnThePet :one
SELECT 1 FROM missing_pets WHERE id = ? AND owner_id = ?;

-- name: UploadPhoto :one
INSERT INTO
missing_pet_photos (pet_id, encoded_data)
VALUES (?, ?)
RETURNING id;

-- name: FindUserByEmail :one
SELECT id, name, phone, email, address, hash FROM pet_owners WHERE email = ?;

-- name: FindUserById :one
SELECT id, name, phone, email, address FROM pet_owners WHERE id = ?;

-- name: CreateUser :one
INSERT INTO
pet_owners (name, phone, email, address, hash)
VALUES (?, ?, ?, ?, ?)
RETURNING id, name, phone, email, address;

-- name: GetUserPets :many
SELECT id, name, type, last_seen FROM missing_pets WHERE owner_id = ?;

-- name: RemoveUserPet :exec
DELETE FROM missing_pets WHERE id = ? AND owner_id = ?;

-- name: GetPetByOwnerAndId :one
SELECT id, name, type, last_seen FROM missing_pets WHERE id = ? AND owner_id = ?;

-- name: GetAllPetsNameFilter :many
SELECT missing_pets.id, missing_pets.name, missing_pets.type, missing_pets.last_seen, pet_owners.id as owner_id
FROM missing_pets 
JOIN pet_owners ON missing_pets.owner_id = pet_owners.id
WHERE (missing_pets.name LIKE CAST(sqlc.arg(name) AS TEXT))
LIMIT sqlc.arg(limit) OFFSET sqlc.arg(offset);
