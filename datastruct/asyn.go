package datastruct

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ProducerAndConsumer ...
func ProducerAndConsumer() {
	cur, max := 0, 10
	mt := &sync.Mutex{}
	produceOne := sync.NewCond(mt)
	consumeOne := sync.NewCond(mt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Producer
	go func() {
		for {
			time.Sleep(3 * time.Second)
			fmt.Println("producer")
			select {
			case <-ctx.Done():
				return
			default:
				mt.Lock()
				defer mt.Unlock()
				for cur == max {
					consumeOne.Wait()
				}
				cur++
				produceOne.Signal()
			}
		}
	}()

	//Consumer
	go func() {

		for {
			time.Sleep(3 * time.Second)
			fmt.Println("consumer")
			select {
			case <-ctx.Done():
				return
			default:
				mt.Lock()
				defer mt.Unlock()
				for cur == 0 {
					produceOne.Wait()
				}
				cur--
				consumeOne.Signal()
			}
		}
	}()

	// fmt.Println(`Input "stop" to exit...`)
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	if input.Text() == "stop" {
	// 		break
	// 	}
	// }
}
