package busy

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Busy(ctx context.Context, wg *sync.WaitGroup) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context Done")
			wg.Done()
			return nil
		default:
			fmt.Println("Sleeping")
			time.Sleep(time.Second * 2)
		}
	}
}
