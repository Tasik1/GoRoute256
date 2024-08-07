package main

import (
	"container/heap"
	"fmt"
	"regexp"
	"strings"
)

type Word struct {
	name     string
	priority int
	index    int
}

type WordQueue []*Word

func (q *WordQueue) Len() int {
	return len(*q)
}

func (q WordQueue) Less(a, b int) bool {
	if q[a].priority == q[b].priority {
		return q[a].name < q[b].name
	}
	return q[a].priority > q[b].priority
}

func (q WordQueue) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
	q[a].index = b
	q[b].index = a
}

func (q *WordQueue) Push(w interface{}) {
	word := w.(*Word)
	word.index = q.Len()
	*q = append(*q, word)
}

func (q *WordQueue) Pop() interface{} {
	currentQueue := *q
	n := q.Len()
	word := currentQueue[n-1]
	word.index = -1
	*q = currentQueue[:n-1]
	return *word
}

func main() {
	pq := make(WordQueue, 0)
	heap.Init(&pq)

	text := "cat and dog, one dog,two cats and one man- какой-то какойто нога Нога нога! нога, нога' hello world"
	text = regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9\-]+`).ReplaceAllString(text, " ")
	text = strings.ToLower(text)
	fmt.Println(text)
	arr := strings.Split(text, " ")
	fmt.Println(arr)
	mp := make(map[string]int)
	for _, val := range arr {
		if val == "" {
			continue
		}
		mp[val]++
	}
	for key, val := range mp {
		heap.Push(&pq, &Word{
			name:     key,
			priority: val,
		})
	}
	for i := 0; i < 10; i++ {
		if pq.Len() == 0 {
			return
		}
		res := heap.Pop(&pq)
		fmt.Println(res)
	}
}
