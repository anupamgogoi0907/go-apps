package test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestContext(t *testing.T) {
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
