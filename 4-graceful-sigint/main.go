//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"os"
	"os/signal"
	"sync/atomic"
)

var callCount int32 = 0

func main() {
	signals := make(chan os.Signal, 1)
	// Create a process
	proc := MockProcess{}

	signal.Notify(signals, os.Interrupt)

	go func() {
		for range signals {
			if atomic.LoadInt32(&callCount) == 0 {
				go proc.Stop()
				atomic.AddInt32(&callCount, 1)
			} else {
				os.Exit(0)
			}
		}
	}()

	// Run the process (blocking)
	proc.Run()
}
