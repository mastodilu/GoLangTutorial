package person

import (
	"fmt"
	"testing"
)

func TestSpeak(t *testing.T) {
	p := Person{
		Firstname: "Matteo",
		Lastname:  "Dilu",
	}

	expected := "I'm Dilu Matteo"
	if got := p.Speak(); got != expected {
		t.Errorf("\nGot '%s' expected '%s'\n", got, expected)
	}
}

func ExampleSpeak() {
	p := Person{
		Firstname: "Matteo",
		Lastname:  "Dilu",
	}

	fmt.Println(p.Speak())
	// Output:
	// I'm Dilu Matteo
}

func BenchmarkSpeak(b *testing.B) {
	p := Person{
		Firstname: "Matteo",
		Lastname:  "Dilu",
	}
	for i := 0; i < b.N; i++ {
		p.Speak()
	}
}
