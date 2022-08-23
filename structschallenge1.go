package main

import "fmt"

type astro struct {
	name string
	age int
	mission string
	isneeded bool
}

func main() {

	astro1 := astro{"Pulkit Chauhan", 22, "Fix ISS", true}
	astro2 := astro{"Brady Collis", 22, "Refuel", false}
	astro3 := astro{"Aidan Gorny", 22, "Mission Control", true}

	fmt.Println(astro1)
	fmt.Println(astro2)
	fmt.Println(astro3)

}