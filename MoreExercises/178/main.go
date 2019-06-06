package main

import (
	"fmt"
)

type CustomErr struct {
	Msg string
}

func (ce CustomErr) Error() string {
	return ce.Msg
}

func main() {
	err := func() error {
		return CustomErr{"ciaone"}
	}()
	if err != nil {
		fmt.Println(err.Error())
	}
}
