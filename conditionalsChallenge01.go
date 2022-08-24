package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {

    rand.Seed(time.Now().Unix())
}

func main() {

    friendsSlice := []string{"Sid", "Brady", "Brian", "Aidan", "Anya", "Sam"}

    if x:=friendsSlice[rand.Intn(len(friendsSlice))]; x == "Sam" {
        fmt.Println("Sam was selected")
    } else if x == "Anya" {
        fmt.Println("Anya was selected")
    } else {
        fmt.Println("Hmmmm, who is", x)
    }

}

