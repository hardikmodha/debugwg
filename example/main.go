package main

import (
	"os"
	"time"

	"github.com/hardikmodha/debugwg"
)

func main() {
	dwg := debugwg.New()

	// Setup periodic print after every second
	stopPeriodicDebug := dwg.PeriodicDebug(os.Stdout, 1 * time.Second)

	dwg.Add(1)

	go func() {
		defer dwg.Done()
		var i int32 = 0

		for i = 0; i< 10; i++ {
			time.Sleep(1 * time.Second)
			dwg.Add(1)
		}

		for i = 0; i< 10; i++ {
			time.Sleep(1 * time.Second)
			dwg.Done()
		}
	}()

	dwg.Wait()

	stopPeriodicDebug()
}