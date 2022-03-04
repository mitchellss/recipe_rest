package storage

import "github.com/segmentio/ksuid"

func GenerateID() (string, error) {
	id, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
