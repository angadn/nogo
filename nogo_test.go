package nogo_test

import (
	"sync"
	"testing"
	"time"

	"github.com/angadn/nogo"
)

func TestBlockAlreadyDone(t *testing.T) {
	block := nogo.NewBlock()
	block.Done()
	block.Wait()
	block.Wait()
	block.Wait()
}

func TestBlockMultipleWait(t *testing.T) {
	block := nogo.NewBlock()
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			<-block.Wait()
			wg.Done()
		}()
	}

	time.Sleep(1 * time.Second)
	block.Done()
	wg.Wait()
	block.Wait()
	block.Done()
}
