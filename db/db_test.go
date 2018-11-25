package db_test

import (
	"testing"

	"github.com/loitd/gesa/db"
)

/*
	Hello is a public function of this test package
	It return just a hello message
*/
func TestHello(t *testing.T) {
	db.Hello("lll")
	// Output:
	// hello from gesa/hello lll
}

func TestConnect(t *testing.T) {
	db.Connect("abc")
	// Output
	// Init OK
}
