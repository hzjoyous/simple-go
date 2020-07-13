package command

import "testing"

func TestDate(t *testing.T) {
	t.Log(date())
}
func BenchmarkDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		date()
	}
}
