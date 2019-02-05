package nogo

import (
	"sync"

	"github.com/google/uuid"
)

type Block struct {
	name string
	done chan struct{}

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
		done: make(chan struct{}),
	}

	return
}

func (b *Block) Wait() {
	<-b.done
}

func (b *Block) Done() {
	b.once.Do(func() {
		close(b.done)
	})
}
