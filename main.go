package main

import (
    "fmt"

    "github.com/rails-to-cosmos/ec2"
)

func main() {
    message := ec2.Hello("Gladys")
    fmt.Println(message)
}
