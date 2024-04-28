package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("missing name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
func randomFormat() string {
	formats := []string{
		"Hi, %v, Welcome",
		"Yo, %v, mofo",
		"Oy, %v ",
	}
	return formats[rand.Intn(len(formats))]

}
