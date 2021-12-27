package main

import (
    "fmt"

    "github.com/karinahallberg/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Karina")
    fmt.Println(message)
}