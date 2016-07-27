// zheng-ji.info

package cuckoo

import (
	"math/rand"
	"sync"
)

// Filter struct
type Filter struct {
	num     int
	buckets []Bucket
	lock    *sync.Mutex
}

// NewFilter Init a Filter with capacity
func NewFilter(capacity uint) *Filter {
	capacity = getCeilingCap(uint64(capacity)) / SlotSize
	if capacity == 0 {
		capacity = 1
	}
	buckets := make([]Bucket, capacity, capacity)
	for i := range buckets {
		buckets[i] = [SlotSize]Signature{}
	}
	return &Filter{
		buckets: buckets,
		num:     0,
		lock:    new(sync.Mutex),
	}
}

// Find Func，check an entry exist or not
func (filter *Filter) Find(data []byte) bool {
	sign := genSignature(data)
	firstIndex := genFirstIndex(sign, uint(len(filter.buckets)))
	backupIndex := genBackupIndex(sign, uint(len(filter.buckets)))

	bk1 := &filter.buckets[firstIndex]
	bk2 := &filter.buckets[backupIndex]

	if bk1.lookupIndex(sign) != NotFound || bk2.lookupIndex(sign) != NotFound {
		return true
	}
	return false
}

// Insert Func，Insert an entry
func (filter *Filter) Insert(data []byte) bool {
	filter.lock.Lock()
	defer filter.lock.Unlock()

	sign := genSignature(data)
	firstIndex := genFirstIndex(sign, uint(len(filter.buckets)))
	backupIndex := genBackupIndex(sign, uint(len(filter.buckets)))
	bk1 := &filter.buckets[firstIndex]
	bk2 := &filter.buckets[backupIndex]
	if bk1.insert(sign) || bk2.insert(sign) {
		filter.num++
		return true
	}
	return filter.resolveCollision(sign, backupIndex)
}

func (filter *Filter) resolveCollision(sign Signature, index uint) bool {
	for i := 0; i < MaxCuckooCount; i++ {
		j := rand.Intn(SlotSize)
		tmpsign := sign
		sign = filter.buckets[index][j]
		filter.buckets[index][j] = tmpsign
		index = genBackupIndex(sign, uint(len(filter.buckets)))
		bk := &filter.buckets[index]
		if bk.insert(sign) {
			filter.num++
			return true
		}
	}
	return false
}

// Del Func:delete entry
func (filter *Filter) Del(data []byte) bool {
	filter.lock.Lock()
	defer filter.lock.Unlock()

	sign := genSignature(data)
	firstIndex := genFirstIndex(sign, uint(len(filter.buckets)))
	backupIndex := genBackupIndex(sign, uint(len(filter.buckets)))
	bk1 := &filter.buckets[firstIndex]
	bk2 := &filter.buckets[backupIndex]
	return bk1.del(sign) || bk2.del(sign)
}

// Size Fun: get size of Filter's element
func (filter *Filter) Size() int {
	return filter.num
}
