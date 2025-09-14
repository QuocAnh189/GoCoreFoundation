package uuid

import "github.com/google/uuid"

func GenerateUUIDV7() (string, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return uuidV7.String(), nil
}
