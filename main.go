package main

import (
	"github.com/lght/systemsonde/data"
	"time"
)

func main() {
	data.Print()
	time.Sleep(4 * time.Second)
	data.Print()
}
