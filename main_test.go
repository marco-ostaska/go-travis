package main

import "testing"

func TestMain(t *testing.T) {

	tt := []struct {
		name   string
		result string
	}{
		{"Not a nice test", "this is a nice test"},
		{"dumb test", "dumb test"},
	}

	for _, tc := range tt {
		if tc.name != tc.result {
			if tc.name == "Not a nice test" {
				t.Skip(tc.name)
			}
			t.Errorf(tc.name)

		}
	}

}
