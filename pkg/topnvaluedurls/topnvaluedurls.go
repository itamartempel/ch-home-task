package topnvaluedurls

import (
	"container/heap"
)

// TopNValuedUrls represent a top list of valued urls with the a size of `topN`
type TopNValuedUrls struct {
	buffer *Buffer
	topN   int
}

func NewTopNValuedUrls(topN int) *TopNValuedUrls {
	ttu := &TopNValuedUrls{
		buffer: &Buffer{&UrlValueTuple{}},
		topN:   topN,
	}
	heap.Init(ttu.buffer)
	return ttu
}

// AddValuedUrl add new valued url to the top list
// in case that the valued url belong to the top list, will push the url in the right place in the list
// in case that the length of the list is bigger than `topN`, will pop the smallest valued url while keeping the order of the list
func (ttu *TopNValuedUrls) AddValuedUrl(url string, value int64) {
	tuple := &UrlValueTuple{
		Url:   url,
		Value: value,
	}
	heap.Push(ttu.buffer, tuple)

	if ttu.buffer.Len() > ttu.topN {
		heap.Pop(ttu.buffer)
	}
}

// GetTopNValuedUrls returns an ordered list of valued urls when the first url with the biggest value and so on with a desending order
func (ttu *TopNValuedUrls) GetTopNValuedUrls() []UrlValueTuple {
	result := make([]UrlValueTuple, ttu.buffer.Len())
	for ttu.buffer.Len() > 0 {
		item := heap.Pop(ttu.buffer).(*UrlValueTuple)

		// ignoring the initial item
		if item.Url != "" {
			reversIndex := ttu.buffer.Len() - 1 - item.Index
			result[reversIndex] = *item
		}
	}
	return result
}

// GetTopNUrls GetTopNValuedUrls returns an ordered list of url strings when the first url with the biggest value and so on with a desending order
func (ttu *TopNValuedUrls) GetTopNUrls() (topNUrls []string) {
	for _, item := range ttu.GetTopNValuedUrls() {
		topNUrls = append(topNUrls, item.Url)
	}
	return
}
