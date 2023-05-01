package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%s %d %d %d, %d:%d:%d\n", now.Weekday(), now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second())
}
