package action

import "testing"

func TestCalc(t *testing.T) {
	res := AddInt(10, 20)
	expectedRes := 30
	if res != expectedRes {
		t.Errorf("10 + 20 = %d", expectedRes)
	}
}

func BenchmarkCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddInt(10, 20)
	}
}
