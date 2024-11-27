package main

import (
	"RyuPlan/go-REST/internal"
	"fmt"
)

func main() {
	fmt.Print("Hello world from outside\n")
	internal.RunRouter()
	return
}
