package queue

import "sync"

type Queue[K any] struct {
	data []K
	mux  sync.Mutex
}

func NewQueue[K any]() *Queue[K] {
	return &Queue[K]{
		data: make([]K, 0),
	}
}

func (q *Queue[K]) Add(i K) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.data = append(q.data, i)
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

func (q *Queue[K]) Empty() bool {
	return len(q.data) <= 0
}
