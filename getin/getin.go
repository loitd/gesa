/*
This work is licensed under a Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.
https://creativecommons.org/licenses/by-nc-sa/4.0/
*/

// Based on tutorial: https://tutorialedge.net/golang/parsing-json-with-golang/#the-encoding-json-package
package getin

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Users struct which contains
// an array of users
type Config struct {
	CONNSTR string `json:"CONNSTR"`
}

// load the config for me
func load() {
	jsonfile, err := os.Open("src/github.com/loitd/gesa/config/main.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonfile.Close()
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	} else {
		// read our opened xmlFile as a byte array.
		byteval, _ := ioutil.ReadAll(jsonfile)
		// we initialize our Users array
		var config Config
		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		json.Unmarshal(byteval, &config)
		// we iterate through every user within our users array and
		// print out the user Type, their name, and their facebook url
		// as just an example
		log.Println(config.CONNSTR)
	}
}

func main() {
	load()
}
