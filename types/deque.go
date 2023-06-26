package types

import "errors"

type Deque[T any] struct {
	Buf    []T
	Head   int
	Tail   int
	Count  int
	MinCap int
}

func New[T any](size ...int) *Deque[T] {
	var cap, min int
	if len(size) >= 1 {
		cap = size[0]
		if len(size) >= 2 {
			min = size[1]
		}
	}

	minCap := 16
	for minCap < min {
		minCap <<= 1
	}

	var buf []T
	if cap != 0 {
		bufSize := minCap
		for bufSize < cap {
			bufSize <<= 1
		}
		buf = make([]T, bufSize)
	}

	return &Deque[T]{
		Buf:    buf,
		MinCap: minCap,
	}
}

func (q *Deque[T]) prev(i int) int {
	return (i - 1) & (len(q.Buf) - 1)
}

func (q *Deque[T]) next(i int) int {
	return (i + 1) & (len(q.Buf) - 1)
}

func (q *Deque[T]) resizeIfFull() {
	if q.Count != len(q.Buf) {
		return
	}
	if len(q.Buf) == 0 {
		if q.MinCap == 0 {
			q.MinCap = 16
		}
		q.Buf = make([]T, q.MinCap)
		return
	}
	q.resize()
}

func (q *Deque[T]) resize() {
	newBuf := make([]T, q.Count<<1)
	if q.Tail > q.Head {
		copy(newBuf, q.Buf[q.Head:q.Tail])
	} else {
		n := copy(newBuf, q.Buf[q.Head:])
		copy(newBuf[n:], q.Buf[:q.Tail])
	}

	q.Head = 0
	q.Tail = q.Count
	q.Buf = newBuf
}

func (q *Deque[T]) shrinkIfExcess() {
	if len(q.Buf) > q.MinCap && (q.Count<<2) == len(q.Buf) {
		q.resize()
	}
}

func outOfRangeText(i, len int) error {
	return errors.New("deque: index out of range " + string(i) + " with length " + string(len))
}

func (q *Deque[T]) Cap() int {
	if q == nil {
		return 0
	}

	return len(q.Buf)
}

func (q *Deque[T]) Len() int {
	if q == nil {
		return 0
	}

	return q.Count
}

func (q *Deque[T]) PushBack(el T) {
	q.resizeIfFull()

	q.Buf[q.Tail] = el
	q.Tail = q.next(q.Tail)
	q.Count++
}

func (q *Deque[T]) PushFront(el T) {
	q.resizeIfFull()

	q.Head = q.prev(q.Head)
	q.Buf[q.Head] = el
	q.Count++
}

func (q *Deque[T]) PopFront() T {
	if q.Count <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	ret := q.Buf[q.Head]
	var zero T
	q.Buf[q.Head] = zero
	q.Head = q.next(q.Head)
	q.Count--

	q.shrinkIfExcess()

	return ret
}

func (q *Deque[T]) PopBack() T {
	if q.Count <= 0 {
		panic("deque: PopBack() called on empty queue")
	}

	q.Tail = q.prev(q.Tail)

	ret := q.Buf[q.Tail]
	var zero T
	q.Buf[q.Tail] = zero
	q.Count--

	q.shrinkIfExcess()

	return ret
}

func (q *Deque[T]) Front() T {
	if q.Count <= 0 {
		panic("deque: Front() called when empty")
	}

	return q.Buf[q.Head]
}

func (q *Deque[T]) Back() T {
	if q.Count <= 0 {
		panic("deque: Back() called when empty")
	}

	return q.Buf[q.prev(q.Tail)]
}

func (q *Deque[T]) At(i int) T {
	if i < 0 || i >= q.Count {
		panic(outOfRangeText(i, q.Len()))
	}

	return q.Buf[(q.Head+i)&(len(q.Buf)-1)]
}

func (q *Deque[T]) Set(i int, item T) {
	if i < 0 || i >= q.Count {
		panic(outOfRangeText(i, q.Len()))
	}

	q.Buf[(q.Head+i)&(len(q.Buf)-1)] = item
}

func (q *Deque[T]) Clear() {
	var zero T
	modBits := len(q.Buf) - 1
	h := q.Head

	for i := 0; i < q.Len(); i++ {
		q.Buf[(h+i)&modBits] = zero
	}

	q.Head = 0
	q.Tail = 0
	q.Count = 0
}

func (q *Deque[T]) Rotate(n int) {
	if q.Len() <= 1 {
		return
	}

	n %= q.Count
	if n == 0 {
		return
	}

	modBits := len(q.Buf) - 1
	if q.Head == q.Tail {
		q.Head = (q.Head + n) & modBits
		q.Tail = q.Head
		return
	}

	var zero T

	if n < 0 {
		for ; n < 0; n++ {
			q.Head = (q.Head - 1) & modBits
			q.Tail = (q.Tail - 1) & modBits
			q.Buf[q.Head] = q.Buf[q.Tail]
			q.Buf[q.Tail] = zero
		}
		return
	}

	for ; n > 0; n-- {
		q.Buf[q.Tail] = q.Buf[q.Head]
		q.Buf[q.Head] = zero
		q.Head = (q.Head + 1) & modBits
		q.Tail = (q.Tail + 1) & modBits
	}
}

func (q *Deque[T]) Index(f func(T) bool) int {
	if q.Len() > 0 {
		modBits := len(q.Buf) - 1
		for i := 0; i < q.Count; i++ {
			if f(q.Buf[(q.Head+i)&modBits]) {
				return i
			}
		}
	}
	return -1
}

func (q *Deque[T]) Insert(at int, item T) {
	if at < 0 || at > q.Count {
		panic(outOfRangeText(at, q.Len()))
	}

	if at*2 < q.Count {
		q.PushFront(item)
		front := q.Head
		for i := 0; i < at; i++ {
			next := q.next(front)
			q.Buf[front], q.Buf[next] = q.Buf[next], q.Buf[front]
			front = next
		}
		return
	}

	swaps := q.Count - at
	q.PushBack(item)
	back := q.prev(q.Tail)

	for i := 0; i < swaps; i++ {
		prev := q.prev(back)
		q.Buf[back], q.Buf[prev] = q.Buf[prev], q.Buf[back]
		back = prev
	}
}

func (q *Deque[T]) Remove(at int) T {
	if at < 0 || at >= q.Len() {
		panic(outOfRangeText(at, q.Len()))
	}

	rm := (q.Head + at) & (len(q.Buf) - 1)
	if at*2 < q.Count {
		for i := 0; i < at; i++ {
			prev := q.prev(rm)
			q.Buf[prev], q.Buf[rm] = q.Buf[rm], q.Buf[prev]
			rm = prev
		}
		return q.PopFront()
	}

	swaps := q.Count - at - 1
	for i := 0; i < swaps; i++ {
		next := q.next(rm)
		q.Buf[rm], q.Buf[next] = q.Buf[next], q.Buf[rm]
		rm = next
	}

	return q.PopBack()
}

func (q *Deque[T]) SetMinCapacity(minCap uint) {
	if 1<<minCap > 16 {
		q.MinCap = 1 << minCap
	} else {
		q.MinCap = 16
	}
}
