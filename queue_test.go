package queue

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestMain(m *testing.M) {
	m.Run()
}

func TestQueue_NewQueue(t *testing.T) {
	q := NewQueue[string]()
	assert.Equal(t, 0, len(q.data), "the len(data) of new queue should be 0")
	assert.Equal(t, true, q.Empty(), "new queue should be empty")
}

func TestQueue_Add(t *testing.T) {
	q := NewQueue[int]()
	data := []int{0, 1, 2}
	for _, n := range data {
		q.Add(n)
	}

	all := q.PopAll()
	assert.Equal(t, len(all), len(data), "length of queue should equal")
	for i, n := range data {
		assert.Equal(t, all[i], n, "item should be equal")
	}
}

func TestQueue_AddList(t *testing.T) {
	q := NewQueue[int]()
	data := []int{0, 1, 2}
	q.AddList(data)

	all := q.PopAll()
	assert.Equal(t, len(all), len(data), "length of queue should equal")
	for i, n := range data {
		assert.Equal(t, all[i], n, "item should be equal")
	}
}

func TestQueue_PopAll(t *testing.T) {
	q := NewQueue[bool]()
	all := q.PopAll()
	assert.Equal(t, 0, len(all))

	q.Add(true)
	all = q.PopAll()
	assert.Equal(t, true, all[0])
}

func TestQueue_Empty(t *testing.T) {
	q := NewQueue[bool]()
	assert.Equal(t, true, q.Empty())

	q.Add(true)
	assert.Equal(t, false, q.Empty())
}
