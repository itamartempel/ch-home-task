package topnvaluedurls

// Buffer is an alias to slice of UrlValueTuple that implement the `heap.Interface` interface
// with the logic of PriorityQueue
type Buffer []*UrlValueTuple

func (pq Buffer) Len() int { return len(pq) }

func (pq Buffer) Less(i, j int) bool {
	return pq[i].Value < pq[j].Value
}

func (pq Buffer) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *Buffer) Push(x interface{}) {
	n := len(*pq)
	item := x.(*UrlValueTuple)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *Buffer) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
