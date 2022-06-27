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
	assert.Equal(t, data, all)
}

func TestQueue_Head(t *testing.T) {
	q := NewQueue[int]()
	q.AddList([]int{0, 1, 2})

	head, err := q.Head()
	assert.Equal(t, 0, head)
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{1, 2}, q.data)

	head, err = q.Head()
	assert.Equal(t, 1, head)
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{2}, q.data)

	head, err = q.Head()
	assert.Equal(t, 2, head)
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{}, q.data)

	head, err = q.Head()
	assert.Equal(t, 0, head)
	assert.Equal(t, ErrEmptyQueue, err)
	assert.Equal(t, []int{}, q.data)
}

func TestQueue_Remove(t *testing.T) {
	q := NewQueue[int]()
	data := []int{0, 1, 2, 3, 4, 5, 3, 9, 8, 6}
	q.AddList(data)

	q.Remove(11)
	shouldBe := []int{0, 1, 2, 3, 4, 5, 3, 9, 8, 6}
	assert.Equal(t, shouldBe, q.data)

	q.Remove(2)
	shouldBe = []int{0, 1, 6, 3, 4, 5, 3, 9, 8}
	assert.Equal(t, shouldBe, q.data)

	q.Remove(3)
	shouldBe = []int{0, 1, 6, 8, 4, 5, 9}
	assert.Equal(t, shouldBe, q.data)

	q.Remove(0)
	shouldBe = []int{9, 1, 6, 8, 4, 5}
	assert.Equal(t, shouldBe, q.data)

	q.Remove(5)
	shouldBe = []int{9, 1, 6, 8, 4}
	assert.Equal(t, shouldBe, q.data)
}

func TestQueue_AddList(t *testing.T) {
	q := NewQueue[int]()
	data := []int{0, 1, 2}
	q.AddList(data)

	all := q.PopAll()
	assert.Equal(t, all, data)
}

func TestQueue_PopAll(t *testing.T) {
	q := NewQueue[bool]()
	all := q.PopAll()
	assert.Equal(t, 0, len(all))

	q.Add(true)
	shouldBe := []bool{true}
	all = q.PopAll()
	assert.Equal(t, shouldBe, all)
}

func TestQueue_Has(t *testing.T) {
	q := NewQueue[int]()
	assert.Equal(t, false, q.Has(5))

	q.Add(9)
	assert.Equal(t, true, q.Has(9))
}

func TestQueue_Each(t *testing.T) {
	q := NewQueue[int]()
	q.AddList([]int{1, 2, 3, 4, 5})
	sum := 0
	q.Each(func(item int) {
		sum = sum + item
	})
	assert.Equal(t, 15, sum)
}

func TestQueue_Empty(t *testing.T) {
	q := NewQueue[bool]()
	assert.Equal(t, true, q.Empty())

	q.Add(true)
	assert.Equal(t, false, q.Empty())
}
