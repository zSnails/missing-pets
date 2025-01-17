-- name: CreateMissingPet :one
INSERT INTO
missing_pets (name, type, last_seen)
VALUES (?, ?, ?)
RETURNING *;

-- name: FindMissingPetsByName :many
SELECT * FROM
missing_pets
WHERE remove_special_characters(name)
LIKE remove_special_characters(CAST(sqlc.arg(name) AS TEXT));

-- name: LinkPetAndOwner :one
INSERT INTO
missing_pet_owner_rel (missing_pet_id, pet_owner_id)
VALUES (?, ?)
RETURNING *;

-- name: UploadPhoto :one
INSERT INTO
missing_pet_photos (pet_id, encoded_data)
VALUES (?, ?)
RETURNING id;
