/*
This work is licensed under a Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.
https://creativecommons.org/licenses/by-nc-sa/4.0/
*/

/*
	This package does anything with database for this app/plugin.
*/
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"strings"

	_ "gopkg.in/rana/ora.v4"
	ora "gopkg.in/rana/ora.v4"
)

/*
	DBStruct defines DB Class
	김치 한국사람
*/
type DBStruct struct {
	// we'll define common properties here
	dbptr *sql.DB
	err   error
	env   *ora.Env
	srv   *ora.Srv
	ses   *ora.Ses
}

/*
	Interfaces for Database Connect Struct
*/
type DBInterfaces interface {
	// connect methods return void
	SimpleConnect()
	Connect()
}

/*
	Hello is a public function of this test package
	It return just a hello message
*/
func Hello(name string) {
	fmt.Println("hello from gesa/hello", name)
}

func (dbins *DBStruct) Connect(constr string) bool {
	dbins.env, dbins.srv, dbins.ses, dbins.err = ora.NewEnvSrvSes(constr)

	defer dbins.env.Close()
	defer dbins.srv.Close()
	defer dbins.ses.Close()

	if dbins.err != nil {
		fmt.Println(dbins.err.Error())
		return false
	} else {
		return true
	}

}

/*
	To connect to database and execute some queries if needed
	Please note:
	- unsupported in Golang: x.f, err := f()
*/
func (dbins *DBStruct) SimpleConnect(constr string) {
	//fmt.Println("Begin Connect function")
	dbins.dbptr, dbins.err = sql.Open("ora", constr)

	//fmt.Println(reflect.TypeOf(db))
	if dbins.err != nil {
		fmt.Println("Err while init: ", dbins.err)
	} else {
		fmt.Println("Init OK")
	}

	defer dbins.dbptr.Close()

	sql1 := "SELECT * FROM DUAL"
	var row string
	rows, err := dbins.dbptr.Query(sql1)

	if err != nil {
		if strings.Contains(err.Error(), "ORA-12560: TNS:protocol adapter error") {
			panic("TNS Error")
		} else {
			panic(err.Error())
		}
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
