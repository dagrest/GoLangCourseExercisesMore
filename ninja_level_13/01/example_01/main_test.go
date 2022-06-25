package example_01

import "testing"

func TestMainFunc(t *testing.T) {
	canine, age := MainFunc()
	if canine.age != 70 {
		t.Errorf("Expected 10, was %d", canine.age)
	}
	if canine.name != "Fido" {
		t.Errorf("Expected 10, was %s", canine.name)
	}
	if age != 140 {
		t.Errorf("Expected 10, was %d", age)
	}
}
