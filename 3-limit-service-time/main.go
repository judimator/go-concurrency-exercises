//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync/atomic"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

const MaxDurationSeconds = 10

var ticker = time.Tick(time.Second)

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	done := make(chan bool)

	if u.IsPremium {
		process()
		return true
	}
	go func() {
		process()
		done <- true
	}()

	for {
		select {
		case isDone := <-done:
			if isDone {
				return true
			}
		case <-ticker:
			if atomic.LoadInt64(&u.TimeUsed) >= MaxDurationSeconds {
				return false
			} else {
				atomic.AddInt64(&u.TimeUsed, 1)
			}
		}
	}
}

func main() {
	RunMockServer()
}
