/*
This work is licensed under a Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.
https://creativecommons.org/licenses/by-nc-sa/4.0/
*/

package sendmail_test

import "testing"

func Test_sendmail(t *testing.T) {
	type args struct {
		from    string
		to      string
		subject string
		body    string
		pass    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendmail(tt.args.from, tt.args.to, tt.args.subject, tt.args.body, tt.args.pass)
		})
	}
}
