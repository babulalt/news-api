package main

import (
	"fmt"

	"github.com/berrybytes/sugam/internal/controller"
)

func main() {
	fmt.Println("Server is running ...")
	controller.Run()
}
