package db_test

import "github.com/loitd/gesa/db"

/*
	Hello is a public function of this test package
	It return just a hello message
*/
func ExampleHello() {
	db.Hello("lll")
	// Output:
	// hello from gesa/hello lll
}
