package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	From     string
	To       string
	Amount   int
	IsMarker bool
	Time     time.Time
}

func sendTransaction(tx Transaction, chans map[string]chan Transaction) {
	key := tx.From + tx.To
	ch, ok := chans[key]
	if !ok {
		fmt.Printf("No channel for %s -> %s\n", tx.From, tx.To)
		return
	}
	ch <- tx
}

func process(name string, in1, in2 <-chan Transaction, out1, out2 chan<- Transaction) {
	go func() {
		for {
			select {
			case tx := <-in1:
				fmt.Printf("Process %s received %d from %s\n", name, tx.Amount, tx.From)
			case tx := <-in2:
				fmt.Printf("Process %s received %d from %s\n", name, tx.Amount, tx.From)
			}
		}
	}()
}

//input - number of processes
// - transactions and snapshot activate as a list:
// {to:, from:, type: [MARKER?], amt:, time:}
// how to store snapshot???
//output - snapshot in a storage
// process values {A:500,B:300 ...}
// channel values {AB:[35,42],BA: ...}

func main() {
	fmt.Println("CHANDY LAMPORT ALGO SIM")

	AB := make(chan Transaction)
	BA := make(chan Transaction)
	BC := make(chan Transaction)
	CB := make(chan Transaction)
	AC := make(chan Transaction)
	CA := make(chan Transaction)

	go process("A", BA, CA, AB, AC)
	go process("B", AB, CB, BA, BC)
	go process("C", AC, BC, CA, CB)

	//central transaction list
	transactions := []Transaction{
		{From: "A", To: "B", Amount: 50, Time: time.Now()},
		{From: "B", To: "C", Amount: 30, Time: time.Now()},
		{From: "C", To: "A", Amount: 20, Time: time.Now()},
	}

	chans := map[string]chan Transaction{
		"AB": AB, "BA": BA,
		"BC": BC, "CB": CB,
		"AC": AC, "CA": CA,
	}

	for _, tx := range transactions {
		go sendTransaction(tx, chans)
	}

	//basic communication system

	//routine creation

	//channel creation

	time.Sleep(time.Second)
}
