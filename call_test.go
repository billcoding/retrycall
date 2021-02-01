package main

import (
	"fmt"
	"testing"
)

func TestCall(t *testing.T) {
	Call(func() interface{} {
		fmt.Println("dsfgdsg")
		fmt.Println("dsfgdsg")
		fmt.Println("dsfgdsg")
		fmt.Println("dsfgdsg")
		fmt.Println("dsfgdsg")
		fmt.Println("dsfgdsg")
		fmt.Println("dsfgdsg")
		panic("hi nil")
	})
}
