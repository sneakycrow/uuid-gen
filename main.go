package main

import (
    "fmt"
    "github.com/satori/go.uuid"
)

func main() {
    u1 := uuid.Must(uuid.NewV4())
    fmt.Printf("%s", u1)
}
