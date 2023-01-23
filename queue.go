package queue

import (
	"errors"
	"sync"
)

var ErrEmptyQueue = errors.New("queue is empty")

type Queue[K comparable] struct {
	data []K
	mux  sync.Mutex
}

func NewQueue[K comparable]() *Queue[K] {
	return &Queue[K]{
		data: make([]K, 0),
	}
}

// Add an element to the tail of queue
func (q *Queue[K]) Add(i K) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.data = append(q.data, i)
}

func (q *Queue[K]) PushWithCallback(i K, cb func(K)) {
	q.mux.Lock()
	defer q.mux.Unlock()

	if cb != nil {
		cb(i)
	}
	q.data = append(q.data, i)
}

func (q *Queue[K]) UnShiftListWithCallback(itemList []K, cb func(K)) {
	q.mux.Lock()
	defer q.mux.Unlock()

	if len(itemList) == 0 {
		return
	}
	for _, m := range itemList {
		cb(m)
	}
	q.data = append(itemList, q.data...)
}

// Head return first element from queue
func (q *Queue[K]) Head() (K, error) {
	q.mux.Lock()
	defer q.mux.Unlock()

	var item K
	if len(q.data) == 0 {
		return item, ErrEmptyQueue
	}
	item = q.data[0]
	q.data = q.data[1:]
	return item, nil
}

func (q *Queue[K]) Remove(i K) {
	q.mux.Lock()
	defer q.mux.Unlock()

	l := len(q.data)
	if l < 1 {
		return
	}
	for index, val := range q.data {
		if val == i {
			q.data[index] = q.data[l-1]
			l--
		}
	}
	q.data = q.data[:l]
}

func (q *Queue[K]) AddList(list []K) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.data = append(q.data, list...)
}

func (q *Queue[K]) PopAll() []K {
	q.mux.Lock()
	defer q.mux.Unlock()
	list := q.data
	q.data = make([]K, 0)
	return list
}

func (q *Queue[K]) Has(item K) bool {
	for i := range q.data {
		if q.data[i] == item {
			return true
		}
	}
	return false
}

func (q *Queue[K]) Each(f func(K)) {
	for _, item := range q.data {
		f(item)
	}
}

func (q *Queue[K]) Empty() bool {
	return len(q.data) <= 0
}

func (q *Queue[K]) FindIndex(cb func(K)bool) int {
	for i, d := range q.data {
		if cb(d) {
			return i
		}
	}
	return -1
}