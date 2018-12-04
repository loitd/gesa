/*
This work is licensed under a Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.
https://creativecommons.org/licenses/by-nc-sa/4.0/
*/

// Based on tutorial: https://tutorialedge.net/golang/parsing-json-with-golang/#the-encoding-json-package

package getin

import "testing"

func Test_load(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"XXX"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			load()
		})
	}
}
