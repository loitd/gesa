/*
This work is licensed under a Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.
https://creativecommons.org/licenses/by-nc-sa/4.0/
*/

package logto

import (
	"log"
	"net/smtp"
)

// send email for me
func sendmail(from string, to string, subject string, body string, pass string) {
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("fucked")
	}
}

func main() {
	sendmail("loitd@vnptepay.com.vn", "loitd@vnptepay.com.vn", "ahihi", "ahihi", "utdodeffatgvkeie")
}
