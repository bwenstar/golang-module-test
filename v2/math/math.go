package bta_math

import "errors"

func Add(a, b int) int {
    return a + b
}

func Sub(a, b int) int {
    return a - b
}

func Multi(a, b int) int {
    return a * b
}

func Divide(a, b int) int, error {
    if b == 0 {
        return 0, errors.New("Unable to divide by zero")
    }
    return a / b, nil
}
