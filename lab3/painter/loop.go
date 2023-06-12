package painter

import (
	"image"
	"sync"

	"golang.org/x/exp/shiny/screen"
)

// Receiver отримує текстуру, яка була підготовлена в результаті виконання команд у циелі подій.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop реалізує цикл подій для формування текстури отриманої через виконання операцій отриманих з внутрішньої черги.
type Loop struct {
	Receiver Receiver
	next     screen.Texture // текстура, яка зараз формується
	prev     screen.Texture // текстура, яка була відправленя останнього разу у Receiver
	MsgQueue MessageQueue
	stopCh   chan struct{}
	mutex    sync.Mutex
	cond     *sync.Cond
}

var size = image.Pt(400, 400)

// Start запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)
	l.MsgQueue = MessageQueue{}
	l.stopCh = make(chan struct{})
	l.cond = sync.NewCond(&l.mutex)
	go l.processEvents()
}

func (l *Loop) processEvents() {
	for {
		select {
		case <-l.stopCh:
			return
		default:
			l.mutex.Lock()
			if l.MsgQueue.IsEmpty() {
				l.cond.Wait()
			}
			op := l.MsgQueue.Pull()
			l.mutex.Unlock()

			if op != nil {
				if update := op.Do(l.next); update {
					l.Receiver.Update(l.next)
					l.next, l.prev = l.prev, l.next
				}
			}
		}
	}
}

func (mq *MessageQueue) IsEmpty() bool {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	return len(mq.Queue) == 0
}

// Post додає нову операцію у внутрішню чергу.

func (l *Loop) Post(op Operation) {
	if op != nil {
		l.MsgQueue.Push(op)
		l.mutex.Lock()
		l.cond.Signal()
		l.mutex.Unlock()
	}
}

// StopAndWait сигналізує
func (l *Loop) StopAndWait() {
	l.stopCh <- struct{}{}
	<-l.stopCh
}

type MessageQueue struct {
	Queue []Operation
	mutex sync.Mutex
}

func (mq *MessageQueue) Push(op Operation) {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	mq.Queue = append(mq.Queue, op)
}

func (mq *MessageQueue) Pull() Operation {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()

	if len(mq.Queue) == 0 {
		return nil
	}

	op := mq.Queue[0]
	mq.Queue = mq.Queue[1:]
	return op
}
