package greeting

import "fmt"

func Hello(name string) string {
	m := fmt.Sprintf("Hi, %v, how are you doing?", name)

	return m
}
