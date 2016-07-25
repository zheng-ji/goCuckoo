// zheng-ji.info
package cuckoo

import (
	"testing"
)

func TestCuckoo(t *testing.T) {
	filter := NewCuckooFilter(10)
	t.Log(getCeilingCap(uint64(10)) / SlotSize)

	filter.Insert([]byte("zheng-ji"))
	filter.Insert([]byte("scut"))
	filter.Insert([]byte("coder"))
	filter.Insert([]byte("stupid"))

	t.Log(filter.buckets)
	t.Log(filter.Size())

	if filter.Find([]byte("stupid")) {
		t.Log("exist")
	} else {
		t.Log("Not exist")
	}

	filter.Del([]byte("stupid"))
	if filter.Find([]byte("stupid")) {
		t.Log("exist")
	} else {
		t.Log("Not exist")
	}

	t.Log(filter.buckets)
	t.Log(filter.Size())
}
