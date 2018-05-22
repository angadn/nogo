package nogo

import (
	"sync"

	"github.com/google/uuid"
)

type Block struct {
	name    string
	blocked []chan bool
	isDone  bool

	l *sync.Mutex
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

func NewBlock() Block {
	ret := Block{}
	ret.name = uuid.New().String()
	return ret
}

func (b *Block) Wait() <-chan bool {
	b.l.Lock()
	defer b.l.Unlock()
	ret := make(chan bool, 1)
	if b.isDone {
		ret <- true
		return ret
	}

	b.blocked = append(b.blocked, ret)
	return ret
}

func (b *Block) Done() {
	b.l.Lock()
	defer b.l.Unlock()

	for _, ch := range b.blocked {
		ch <- true
	}

	b.blocked = []chan bool{} // Let those channels get garbage collected
	b.isDone = true
}
