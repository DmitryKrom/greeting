package greeting

import "fmt"

func Hello(name string) string {
	massage := fmt.Sprintf("Hi, %v, how are you doing?", name)

	return massage
}
