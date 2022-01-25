package test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestContext(t *testing.T) {
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	// worker1
	wg.Add(1)
	go func() {
		res, err := worker1(0, ctx, &wg)
		if err != nil {
			cancel()
		}
		fmt.Println(res)
	}()

	// worker2
	wg.Add(1)
	go func() {
		res, _ := worker2(1, ctx, &wg)
		fmt.Println(res)
	}()

	wg.Wait()
}
func worker1(val int, ctx context.Context, wg *sync.WaitGroup) (string, error) {
	if val == 0 {
		wg.Done()
		return "", errors.New("Wrong data")
	} else {
		wg.Done()
		return "worker1 processed.", nil
	}
}
func worker2(val int, ctx context.Context, wg *sync.WaitGroup) (string, error) {
	var res string = ""
	var err error = nil
	select {
	case <-ctx.Done():
		fmt.Println("Worker2 stopped.")
	default:
		if val == 0 {
			wg.Done()
			res = ""
			err = errors.New("Wrong data")
		} else {
			wg.Done()
			res = "worker2 processed"
			err = nil
		}
	}
	return res, err

}
