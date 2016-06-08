package main

import (
	"fmt"

	"github.com/Nitecon/iqfeed"
)

func main() {
	iq := &iqfeed.IQC{BackupFile: "../TestMyStuffs.dat", CreateBackup: true}
	c := iq.Start("127.0.0.1:5009")
	c.ReqAllUpdateFNames()
	c.RequestStats()
	c.RequestWatches()
	c.NewsOn()
	c.WatchSymbol("AAPL1610F90")
	c.WatchSymbol("AAPL1610F90.5")
	c.WatchSymbol("AAPL1610F91")
	c.WatchSymbol("AAPL1610F91.5")
	c.WatchSymbol("AAPL1610F92")
	c.WatchSymbol("AAPL1610F92.5")
	c.WatchSymbol("AAPL1610F93")
	c.WatchSymbol("AAPL1610F93.5")
	c.WatchSymbol("AAPL1610F94")
	c.WatchSymbol("AAPL1610F94.5")
	c.WatchSymbol("AAPL1610F95")
	c.WatchSymbol("AAPL1610F95.5")
	c.WatchSymbol("AAPL1610F96")
	c.WatchSymbol("AAPL1610F96.5")
	c.WatchSymbol("AAPL1610F97")
	c.WatchSymbol("AAPL1610F97.5")
	c.WatchSymbol("AAPL1610F98")
	c.WatchSymbol("AAPL1610F98.5")
	c.WatchSymbol("AAPL1610F99")
	c.WatchSymbol("AAPL1610F99.5")

	c.WatchSymbol("AAPL1610R90")
	c.WatchSymbol("AAPL1610R90.5")
	c.WatchSymbol("AAPL1610R91")
	c.WatchSymbol("AAPL1610R91.5")
	c.WatchSymbol("AAPL1610R92")
	c.WatchSymbol("AAPL1610R92.5")
	c.WatchSymbol("AAPL1610R93")
	c.WatchSymbol("AAPL1610R93.5")
	c.WatchSymbol("AAPL1610R94")
	c.WatchSymbol("AAPL1610R94.5")
	c.WatchSymbol("AAPL1610R95")
	c.WatchSymbol("AAPL1610R95.5")
	c.WatchSymbol("AAPL1610R96")
	c.WatchSymbol("AAPL1610R96.5")
	c.WatchSymbol("AAPL1610R97")
	c.WatchSymbol("AAPL1610R97.5")
	c.WatchSymbol("AAPL1610R98")
	c.WatchSymbol("AAPL1610R98.5")
	c.WatchSymbol("AAPL1610R99")
	c.WatchSymbol("AAPL1610R99.5")

	for {
		select {
		case system := <-c.System:
			fmt.Printf("System: %#v\n", system)
		case fund := <-c.Fundamental:
			fmt.Printf("Fundamental: %#v\n", fund)
		case err := <-c.Errors:
			fmt.Printf("Errors: %#v\n", err)
		case updates := <-c.Updates:
			fmt.Printf("Updates: %#v\n", updates)
		case tm := <-c.Time:
			fmt.Printf("Time: %#v\n", tm)
		case reg := <-c.Regional:
			fmt.Printf("Regional: %#v\n", reg)
		case news := <-c.News:
			fmt.Printf("News: %#v\n", news)
		}

	}
}
