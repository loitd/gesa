/*
This work is licensed under a Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.
https://creativecommons.org/licenses/by-nc-sa/4.0/
*/

/*
	This package does anything with logfiles.
	Currently developed and tested on Windows OS.
*/
package logto

import (
	"log"
	"os"
)

func logto() {
	fh, err := os.OpenFile("src/github.com/loitd/gesa/config/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	log.SetOutput(fh)
	log.Println("yes, locked")
}

func main() {
	logto()
}
