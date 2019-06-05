package rolldice

import (
	"fmt"
	"testing"
)

func TestRollDice(t *testing.T) {
	// expecting numbers between 1 and 6 included
	const iterations = 1000 // number of test to execute
	min, max := 1, 6
	for i := 0; i < iterations; i++ {
		if got := RollDice(); got < min || got > max {
			t.Errorf("Expecting [%d;%d] but got %d.\n", min, max, got)
		}
	}
}

func ExampleRollDice() {
	x := RollDice()
	fmt.Println(x > 0 && x < 7)
	// Output:
	// true
}

func BenchmarkRollDice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RollDice()
	}
}
