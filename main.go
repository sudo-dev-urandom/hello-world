package main

import (
	"fmt"
	"hello-world/controller"
)

func main() {
	count := controller.SearchApi("Avengers", "faf7e5bb&s", 2)
	fmt.Println(count)
	// controller.GetDetail("tt4853102", "faf7e5bb&s", 2)
}
