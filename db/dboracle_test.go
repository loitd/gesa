package db_test

import (
	"testing"

	"github.com/loitd/gesa/db"
)

func TestGetdual(t *testing.T) {
	dbi := db.DBStruct()
	dbi.Getdual()

}

// func TestConnect(t *testing.T) {
// 	dbins := db.DBStruct{}
// 	tabs := []struct {
// 		x string
// 		y string
// 	}{
// 		{"abc/cdx", "ORA-12560: TNS:protocol adapter error"},
// 		{"VIRTUAL_ACCOUNT/VIRTUAL_ACCOUNT123@172.16.10.84:1521/DB84", "Database connected successfully"},
// {"VIRTUAL_ACCOUNT/VIRTUAL_ACCOUNT1234@172.16.10.84:1521/DB84", "ORA-12170: TNS:Connect timeout occurred"},
// {"VIRTUAL_ACCOUNT/xxx@172.16.10.84:1521/DB85", "ORA-12170: TNS:Connect timeout occurred"},
// 	}
// 	for _, tab := range tabs {
// 		err := reflections.SetField(&dbins, "Constr", tab.x)
// 		if err != nil {
// 			t.Error("Unable to set the connection string", err)
// 		}
// 		// dbins.constr = tab.x
// 		res := dbins.Connect()
// 		if !strings.Contains(res, tab.y) {
// 			t.Errorf("[FAILED]in string: %v. result: %v. want: %v", dbins.Constr, res, tab.y)
// 		} else {
// 			t.Logf("[PASSED]in string: %v. result: %v. want: %v", dbins.Constr, res, tab.y)
// 		}
// 	}
// }
