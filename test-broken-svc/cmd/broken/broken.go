package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("working....")
	time.Sleep(10 * time.Second)
	log.Fatalf("uh oh...")
}
