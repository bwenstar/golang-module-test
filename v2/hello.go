package hello

import "fmt"

func Hello(name string) string {
    message := fmt.Sprintf("Hello, %v! (v2)", name)
    return message
}

func Goodbye(name string) string {
    message := fmt.Sprintf("Goodbye, %v! (v2)", name)
    return message
}
