/*
This work is licensed under a Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.
https://creativecommons.org/licenses/by-nc-sa/4.0/
*/

// This package does anything with database for this app/plugin.
// Currently added Oracle database via ora.v4 developed and tested on Windows OS.
// Created by Loitd for his project
package db

import (
	"fmt"
	"log"

	"gopkg.in/rana/ora.v4"

	"github.com/oleiade/reflections"
)

/*
	DBStruct defines DB Class
	김치 한국사람
*/
type DBStruct struct {
	// we'll define common properties here
	CONNSTR string
}

// Init all
func (dbi *DBStruct) init() {
	//
	err := reflections.SetField(dbi, "CONNSTR", "VIRTUAL_ACCOUNT/VIRTUAL_ACCOUNT123@172.16.10.84:1521/DB84")
	if err != nil {
		fmt.Println("Unable to set the connection string", err)
	}
}

// get dual
func (dbi *DBStruct) GetDual() string {
	dbi.init()
	var retval string
	env, srv, ses, err := ora.NewEnvSrvSes(dbi.CONNSTR)

	if err != nil {
		log.Fatal(err)
	} else {
		defer env.Close()
		defer srv.Close()
		defer ses.Close()
		// defer after checking error
		stmtQry, err := ses.Prep(fmt.Sprintf("SELECT * FROM %v", "TBL_FUNCTION"))
		if err != nil {
			log.Fatal(err)
		} else {
			defer stmtQry.Close()
			// defer after checking error
			if rset, err := stmtQry.Qry(); err != nil {
				log.Fatal(err)
			} else {
				// for rset.Next() {
				// 	fmt.Println(rset.Row)
				// }
				// if err := rset.Err(); err != nil {
				// 	log.Fatal(err)
				// }
				//-------------------------------------
				row := rset.NextRow()
				if row != nil {
					fmt.Println(row[2])
					retval = fmt.Sprintf("%v", row[2])
				}
				if err := rset.Err(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	return retval
}

// main
func main() {
	dbi := DBStruct{}
	fmt.Println(dbi.GetDual())
}
