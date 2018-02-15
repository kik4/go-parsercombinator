package main

import (
	"fmt"
	ps "go-parsercombinator"
)

func main() {
	fmt.Println(ps.AnyRune().Once().Parse("ABC"))
}
