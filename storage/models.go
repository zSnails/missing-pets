// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package storage

type MissingPet struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	LastSeen string `json:"lastSeen"`
	OwnerID  int64  `json:"ownerId"`
}

type MissingPetPhoto struct {
	ID          int64  `json:"id"`
	PetID       int64  `json:"petId"`
	EncodedData string `json:"encodedData"`
}

type PetOwner struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Hash    []byte `json:"hash"`
}
