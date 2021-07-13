package greeting

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name given")
	}
	m := fmt.Sprintf("Hi, %v, how are you doing?", name)

	return m, nil
}
