package main 

import (
    "fmt"
    "go-client/src"
)
func main() {
    fmt.Println("Welcome to the Go Rest Client!");
    fmt.Println("Choose the entity you wanna see:");
    fmt.Println("----------------------------------")


    choice := src.Input("Choose the entity 1. Users 2. Posts 3. Comments: ")

    fmt.Println("----------------------------------")

    //Swich and fetch the data as per the choice
    switch choice {
    case "1":
        fmt.Println("Fetching users data...")
    case "2":
        fmt.Println("Fetching posts data...")
    case "3":
        fmt.Println("Fetching comments data...")
    default:
        fmt.Println("Invalid choice")
    }
    fmt.Println("Thank you for using the Go Rest Client!")
    


}