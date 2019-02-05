package nogo

import (
	"sync"

	"github.com/google/uuid"
)

type Block struct {
	name string
	wait chan struct{}

	once sync.Once
}

func (b Block) Equals(o Block) bool {
	return b.name == o.name
}

func (b Block) WithName(name string) Block {
	b.name = name
	return b
}

func (b Block) String() string {
	return b.name
}

func NewBlock() (block Block) {
	block = Block{
		name: uuid.New().String(),
		wait: make(chan struct{}),
	}

	return
}

func (b *Block) Wait() (wait <-chan struct{}) {
	wait = b.wait
	return
}

func (b *Block) Done() {
	b.once.Do(func() {
		close(b.wait)
	})
}
