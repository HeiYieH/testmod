package testmod

import "fmt"

// SayHello returns a friendly greeting
func SayHello(name string) string {
    return fmt.Sprintf("Hello, %s", name)
}
