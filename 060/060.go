package main

import (
	"errors"
	"fmt"
	"log"
)

type myerror struct {
	s    string
	a, b int
}

func (m myerror) Error() string { // 💥
	return fmt.Sprintf("error: %s - %d-%d", m.s, m.a, m.b)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("negative number") // 💥
		// return 0, fmt.Errorf("negative number") // 💥
	}
	return 42, nil
}

func attachErrorToMyStruct() (int, error) {
	myerr := myerror{
		s: "my error message",
		a: 88,
		b: 77,
	}
	return 42, myerr
}

func main() {
	fmt.Println("Errors with info")

	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}

	num, err := attachErrorToMyStruct()
	fmt.Println(num, err)
}
