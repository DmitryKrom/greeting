package greeting

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	m := fmt.Sprintf("Hi, %v, how are you doing?", name)
	if name == "" {
		return "", errors.New("no name given")
	}

	return m, nil
}
