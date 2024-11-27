package cmd

import (
	"RyuPlan/go-REST/internal"
	"fmt"
)

func Run() {
	fmt.Print("Hello world from outside\n")
	internal.RunRouter()
	return
}
