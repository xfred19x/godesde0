package main

import (
	"fmt"

	"github.com/xfred19x/godesde0/17.DEFER_PANIC_RECOVER/defer_panic"
)

func main() {

	fmt.Println("-------------")
	defer_panic.VemosDefer()

	fmt.Println("-------------")
	defer_panic.EjemploPanic()
}
