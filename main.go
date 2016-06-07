package main

import (
	"fmt"

	"github.com/Nitecon/iqfeed"
)

func main() {
	iq := &iqfeed.IQC{}
	c := iq.Start("127.0.0.1:5009")
	for {
		select {
		case system := <-c.System:
			fmt.Printf("System Data: %+v\n", system)
		case fund := <-c.Fundamental:
			fmt.Printf("Fundamental Data: %+v\n", fund)
		case err := <-c.Errors:
			fmt.Printf("Errors: %+v\n", err)
		case updates := <-c.Updates:
			fmt.Printf("Updates: %+v\n", updates)
		case tm := <-c.Time:
			fmt.Printf("Time: %+v\n", tm)
		case reg := <-c.Regional:
			fmt.Printf("Regional: %+v\n", reg)
		case news := <-c.News:
			fmt.Printf("News: %+v\n", news)
		}

	}
}
