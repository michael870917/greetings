package grettings

import (
	"errors"
	"fmt"
	"math/rand"
)

// add func

func HelloFromGrettings(name string) (string, error) {

	if name == "" {
		return name, errors.New("Name empty!")
	}

	message := fmt.Sprintf(RandomFormat(), name)

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// Definir map con key y value string
	messages := make(map[string]string)
	for _, name := range names {
		message, err := HelloFromGrettings(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func RandomFormat() string {
	formats := []string{
		"Hello, %v! Welcome here!",
		"Hello, %v! Nice to meet you!",
		"Hello, %v! You are great!",
	}

	return formats[rand.Intn(len(formats))]
}
