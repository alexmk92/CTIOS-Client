package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)

	fmt.Println("I'm the poller");
	c := Connection {
		host: "10.96.134.31",
		port: 42027,
		invokedID: 0,
	}
	c.connect()

	<-done
}
