package topnvaluedurls

import (
	"fmt"
	"math/rand"
	"testing"
)

func createDummyTopTwoValuedUrls() *TopNValuedUrls {
	ttu := NewTopNValuedUrls(2)
	ttu.AddValuedUrl("http://api.tech.com/item/123345", 25)
	ttu.AddValuedUrl("http://api.tech.com/item/124345", 231)
	ttu.AddValuedUrl("http://api.tech.com/item/125345", 111)
	return ttu
}

func TestTopNValuedUrlsAddValuedUrl(t *testing.T) {
	ttu := createDummyTopTwoValuedUrls()
	b := *ttu.buffer

	if b[1].Value != 231 {
		t.Errorf("Expected the second item in the slice to be with the biggest value, get %d", b[1].Value)
	}

	if b[0].Value != 111 {
		t.Errorf("Expected the first item in the slice to be with the second biggest value, get %d", b[0].Value)
	}

	if len(b) != 2 {
		t.Errorf("Expected the buffer slice to be the same length as topN, get %d", len(*ttu.buffer))
	}
}

func TestTopNValuedUrlsGetTopTenValuedUrls(t *testing.T) {
	ttu := createDummyTopTwoValuedUrls()
	a := ttu.GetTopNValuedUrls()

	if a[0].Value != 231 {
		t.Errorf("Expected the first item in the slice to be with the biggest value, get %d", a[0].Value)
	}

	if a[1].Value != 111 {
		t.Errorf("Expected the first item in the slice to be with the second biggest value, get %d", a[1].Value)
	}

	if len(a) != 2 {
		t.Errorf("Expected the slice to be the same length as topN, get %d", len(*ttu.buffer))
	}
}

func benchmarkTopNValuedUrls(i int, b *testing.B) {
	ttu := NewTopNValuedUrls(i)
	for n := 0; n < b.N; n++ {
		ttu.AddValuedUrl(fmt.Sprintf("http://api.tech.com/item/%d", 100_000+rand.Intn(900_000)), rand.Int63n(1000_000))
	}
}

func BenchmarkTopNValuedUrls2(b *testing.B)    { benchmarkTopNValuedUrls(2, b) }
func BenchmarkTopNValuedUrls10(b *testing.B)   { benchmarkTopNValuedUrls(10, b) }
func BenchmarkTopNValuedUrls50(b *testing.B)   { benchmarkTopNValuedUrls(50, b) }
func BenchmarkTopNValuedUrls200(b *testing.B)  { benchmarkTopNValuedUrls(200, b) }
func BenchmarkTopNValuedUrls1000(b *testing.B) { benchmarkTopNValuedUrls(1000, b) }
