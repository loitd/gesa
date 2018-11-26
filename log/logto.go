/*
	This package does anything with logfiles.
	Currently developed and tested on Windows OS.
*/
package log

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("MINGW_EXEC: ", os.Getenv("MINGW_EXEC"))
}
