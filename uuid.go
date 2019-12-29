package helper

import "github.com/google/uuid"

func Uuid() string {
	return uuid.Must(uuid.NewRandom()).String()
}

