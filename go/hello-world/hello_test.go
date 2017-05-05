package hello

import "testing"

const targetTestVersion = 2

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name, expected string
	}{
		{"", "Hello, World!"},
		{"Gopher", "Hello, Gopher!"},
	}
	for _, test := range tests {
		observed := HelloWorld(test.name)
		if observed != test.expected {
			t.Fatalf("HelloWorld(%s) = %v, want %v", test.name, observed, test.expected)
		}
	}

	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}
