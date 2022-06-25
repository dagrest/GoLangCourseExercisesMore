package dog

import "testing"

func TestYears(t *testing.T) {
	age := Years(10)
	if age != 70 {
		t.Errorf("Expected 70, was %d", age)
	}
}

func TestYearsTwo(t *testing.T) {
	age := YearsTwo(20)
	if age != 140 {
		t.Errorf("Expected 140, was %d", age)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(20)
	}
}
