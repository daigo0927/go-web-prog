package main

import (
	"fmt"
	"time"
)

func printNumbers2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
}

func printLetters2() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
}

func print2() {
	printNumbers2()
	printLetters2()
}

func goPrint2() {
	go printNumbers2()
	go printLetters2()
}

func main() {
}
