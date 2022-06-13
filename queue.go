package queue

import "sync"

type Queue[K comparable] struct {
	data []K
	mux  sync.Mutex
}

func NewQueue[K comparable]() *Queue[K] {
	return &Queue[K]{
		data: make([]K, 0),
	}
}

func (q *Queue[K]) Add(i K) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.data = append(q.data, i)
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

func (q *Queue[K]) Each(f func(K)) {
	for _, item := range q.data {
		f(item)
	}
}

func (q *Queue[K]) Empty() bool {
	return len(q.data) <= 0
}
