package topnvaluedurls

import "testing"

func createDummyBuffer() Buffer {
	return Buffer{
		&UrlValueTuple{Url: "http://api.tech.com/item/122345", Value: 350, Index: 0},
		&UrlValueTuple{Url: "http://api.tech.com/item/124345", Value: 231, Index: 1},
		&UrlValueTuple{Url: "http://api.tech.com/item/125345", Value: 111, Index: 2},
		&UrlValueTuple{Url: "http://api.tech.com/item/123345", Value: 25, Index: 3},
		&UrlValueTuple{Url: "http://api.tech.com/item/121345", Value: 9, Index: 4},
	}
}

func TestBufferLess(t *testing.T) {
	b := createDummyBuffer()
	if b.Less(0, 1) {
		t.Error("Less is doing more :)")
	}
}

func TestBufferSwap(t *testing.T) {
	b := createDummyBuffer()
	b.Swap(0, 1)
	if b[0].Value != 231 || b[1].Value != 350 {
		t.Errorf("Expected the fist and the second items to swap places")
	}
}

func TestBufferPush(t *testing.T) {
	b := createDummyBuffer()
	b.Push(&UrlValueTuple{Url: "http://api.tech.com/item/12345", Value: 8})
	if b[len(b)-1].Value != 8 {
		t.Errorf("Expected new item to push to the end")
	}
}

func TestBufferPop(t *testing.T) {
	b := createDummyBuffer()
	popedItem := b.Pop().(*UrlValueTuple)
	if popedItem.Value != 9 || b[len(b)-1].Value == 9 || len(b) != 4 {
		t.Errorf("Expected the last item of the slice to pop out and shrink the slice")
	}
}
