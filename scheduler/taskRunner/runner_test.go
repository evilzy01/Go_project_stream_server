package taskRunner

import (
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChnn) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %v", i)
		}
		return nil
	}

	e := func(dc dataChnn) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("Executer received: %v", d)
			default:
				break forloop
			}
		}
		return nil
	}

	runner := NewRunner(30, false, d, e)
	go runner.startAll()        // goroutine, because startDispatch is a dead loop!
	time.Sleep(time.Second * 3) // while the goroutine is running, another goroutin_main function, runs time.Sleep,
	// that's why the result always stop after about 3 secs.
}
