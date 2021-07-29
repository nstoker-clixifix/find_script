package main

import (
	"fmt"

	d "github.com/nstoker-clixifix/find_script/internal/database"
)

func main() {
	fmt.Println("Hello, world!")

	d.Connect()
}
