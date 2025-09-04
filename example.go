package main

import (
	"fmt"
	"time"
)

func process(name string, in1, in2, out1, out2 chan string) {

	go func() {
		for {
			select {
			case msg := <-in1:
				fmt.Printf("Process %s received (from in1): %s\n", name, msg)
			case msg := <-in2:
				fmt.Printf("Process %s received (from in2): %s\n", name, msg)
			}
		}
	}()

	outmsg := "hello from " + name
	out1 <- outmsg
	out2 <- outmsg

}

func main() {
	fmt.Println("CHANDY LAMPORT ALGO SIM")

	//input - number of processes
	// - transactions and snapshot activate as a list:
	// {to:, from:, type: [MARKER?], amt:, time:}
	// how to store snapshot???
	//output - snapshot in a storage
	// process values {A:500,B:300 ...}
	// channel values {AB:[35,42],BA: ...}

	//basic communication system

	//routine creation

	//channel creation
	AB := make(chan string)
	BA := make(chan string)
	BC := make(chan string)
	CB := make(chan string)
	AC := make(chan string)
	CA := make(chan string)

	go process("A", BA, CA, AB, AC)
	go process("B", AB, CB, BA, BC)
	go process("C", AC, BC, CA, CB)

	time.Sleep(time.Second)
}
