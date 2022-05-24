package main

import (
	"fmt"
	"time"
)

type ball struct {
	hits int
}

func player(name string, table chan *ball) {
	for {
		ball := <- table
		ball.hits++
		fmt.Println("player", name, "hits ball.")
		time.Sleep(time.Millisecond * 500)
		table <- ball
	}
}

func main() {
	table := make(chan *ball)

	go player("ping", table)
	go player("pong", table)

	table <- new(ball)
	time.Sleep(time.Second * 5)

	<-table
}