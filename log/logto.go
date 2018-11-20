package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("MINGW_EXEC: ", os.Getenv("MINGW_EXEC"))
}
