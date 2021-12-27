package main

import (
    "fmt"

    "github.com/karinahallberg/greetings.git"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Karina")
    fmt.Println(message)
}