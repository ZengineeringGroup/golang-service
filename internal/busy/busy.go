package busy

import (
	"context"
	"fmt"
	"time"
)

func Busy(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context Done")
			return nil
		default:
			fmt.Println("Sleeping")
			time.Sleep(time.Second * 2)
		}
	}
}
