# debugwg

Debug Go WaitGroup by Printing WaitGroup count

1. To enable debug, replace the `sync.WaitGroup{}` with `debugwg.New()`
2. Use helper functions `wg.PeriodicDebug` or `wg.Debug` to print the debug information.


```go
package main

import (
 "sync"
 "time"
)

func main() {
 wg := sync.WaitGroup{}
 // wg := debugwg.New()

 // Setup periodic print after every second
 // stopPeriodicDebug := wg.PeriodicDebug(os.Stdout, 1 * time.Second)

 wg.Add(1)

 go func() {
  defer wg.Done()
  var i int32 = 0

  for i = 0; i< 10; i++ {
   time.Sleep(1 * time.Second)
   wg.Add(1)
  }

  for i = 0; i< 9; i++ {
   time.Sleep(1 * time.Second)
   wg.Done()
  }
 }()

 wg.Wait()

 // stopPeriodicDebug()
}
```
