package debugwg

import (
	"fmt"
	"io"
	"sync"
	"time"

	"go.uber.org/atomic"
)

type stopPeriodicDebugFunc func()

type debugWg struct {
	wg    *sync.WaitGroup
	count *atomic.Int32
	stopChan chan struct{}
}

// DebugWg defined methods that are superset of sync.WaitGroup. Extra methods can be used to print debugging
// information.
type DebugWg interface {
	Add(delta int)
	Done()
	Wait()
	Debug(out io.Writer)
	PeriodicDebug(out io.Writer, interval time.Duration) stopPeriodicDebugFunc
}

// New returns the DebugWg interface implementation. It accepts the out
func New() DebugWg {
	return &debugWg{
		wg:    &sync.WaitGroup{},
		count: atomic.NewInt32(0),
		stopChan: make(chan struct{}, 1),
	}
}

func (dwg *debugWg) Add(delta int) {
	dwg.wg.Add(delta)
	dwg.count.Add(int32(delta))
}

func (dwg *debugWg) Done() {
	dwg.wg.Done()
	dwg.count.Add(-1)
}

func (dwg *debugWg) Wait() {
	dwg.wg.Wait()
}


func (dwg *debugWg) Debug(out io.Writer) {
	out.Write([]byte(fmt.Sprintf("DebugWg Count: %d\n", dwg.count.Load())))
}

func (dwg *debugWg) PeriodicDebug(out io.Writer, interval time.Duration) stopPeriodicDebugFunc {
	go func() {
		for {
			select {
			case <-dwg.stopChan:
				return
			case <-time.After(interval):
				dwg.Debug(out)
			}
		}
	}()

	return func() {
		dwg.stopChan <- struct{}{}
	}
}

