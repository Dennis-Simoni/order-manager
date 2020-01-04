package main

import "testing"

var tests = []struct {
	name, expected string
}{
	{"", "Hey you!, welcome to the project!"},
	{"Alice", "Hey Alice, welcome to the project!"},
	{"Bob", "Hey Bob, welcome to the project!"},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		if observed := greeting(test.name); observed != test.expected {
			t.Fatalf("Greeting(%s) = \"%v\", want \"%v\"", test.name, observed, test.expected)
		}
	}
}
