package main

type Queue struct {
	items []*interface{}
}

func (q *Queue) enqueue(item interface{}) {
	q.items = append(q.items, &item)
}

func (q *Queue) dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	first := q.items[0]
	q.items = q.items[1:]
	return *first
}
