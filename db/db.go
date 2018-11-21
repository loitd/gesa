/*
This work is licensed under a Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.
https://creativecommons.org/licenses/by-nc-sa/4.0/
*/

/*
	This package does anything with database for this app/plugin.
*/
package db

import (
	//"database/sql"
	"fmt"
	"log"
	"os/exec"
	//_ "gopkg.in/rana/ora.v4"
)

/*
	Hello is a public function of this test package
	It return just a hello message
*/
func Hello(name string) {
	fmt.Println("hello from gesa/hello", name)
}

/*
func connect(constr string) {
	db, err := sql.Open("ora", constr)

	if err != nil {
		fmt.Println("Err while init: ", err)
	} else {
		fmt.Println("Init successfully ...")
	}

	defer db.Close()

	sql1 := "SELECT * FROM DUAL"
	var row string
	rows, err := db.Query(sql1)

	if err != nil {
		fmt.Println("Err while query: ", err)
	} else {
		fmt.Println("Query successfully...")
		for rows.Next() {
			err = rows.Scan(&row)
			if err != nil {
				fmt.Println("Error while reading: ", err)
				break
			} else {
				fmt.Println("Reading results: ", row)
			}
		}
	}

	defer rows.Close()

}
*/

func ping(target string) string {
	out, err := exec.Command("ping", "8.8.8.8").Output()
	if err != nil {
		log.Fatal(err)
		return string("err")
	} else {
		return string(out)
	}
}

//func main() {
// res := ping("8.8.8.8")
// fmt.Print(res)
//connect("VIRTUAL_ACCOUNT/VIRTUAL_ACCOUNT1234@172.16.10.84:1521/DB84")
//}
