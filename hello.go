package main

// package is a collection fo many related files with .go extension
// The very first line of every file in go tells what package the file belongs to

// In go we have two types of packages one which can be run called the executable and the other is the reusable logic which can be reused (Like helper funciton and stuff)
/*
- But how do we say if a package is executable or reusable
	- The name main in after the keyword package tells go that this is an executable file (Package main ---> go build ---> main)
	- Any name other than main will make the package reuseable (Package main ---> go build ---> nothing)
*/

import "fmt"

func main() {
	fmt.Println("Hello World")
}
