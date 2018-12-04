package logto

import (
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func write() {
	xlsx, err := excelize.OpenFile("src/github.com/loitd/gesa/config/template.xlsx")
	if err != nil {
		log.Fatal(err)
		return
	}
	xlsx.SetCellValue("Sheet1", "A2", "ahihi")
	err = xlsx.SaveAs("src/github.com/loitd/gesa/config/out1.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}
