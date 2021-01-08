//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, buff chan Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(buff)
			return
		}
		buff <- *tweet
	}
}

func consumer(buff chan Tweet) {
	for t := range buff {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	buff := make(chan Tweet)
	// Producer
	go producer(stream, buff)

	// Consumer
	consumer(buff)

	fmt.Printf("Process took %s\n", time.Since(start))
}
