// zheng-ji.info
package cuckoo

import (
	"encoding/binary"
	"hash/fnv"
	"math"
)

func getCeilingCap(capacity uint64) uint {
	num := 1
	for ; capacity/2 != 0; capacity = capacity / 2 {
		num += 1
	}
	return uint(math.Pow(2, float64(num)))
}

func genFp(data []byte) Fp {
	hashInstance := fnv.New64()
	hashInstance.Reset()
	hashInstance.Write(data)
	hash := hashInstance.Sum(nil)
	fp := Fp{}
	for i := 0; i < FingerSize; i++ {
		fp[i] = hash[i]
	}
	if fp == Empty {
		fp[0] ^= 1
	}
	return fp
}

func genFirstIndex(fp Fp, numBuckets uint) uint {
	bytes := make([]byte, 64, 64)
	for i, b := range fp {
		bytes[i] = b
	}
	hash := binary.LittleEndian.Uint64(bytes)
	return uint(hash) & (numBuckets - 1)
}

func genBackupIndex(fp Fp, numBuckets uint) uint {
	bytes := make([]byte, 64, 64)
	for i, b := range fp {
		bytes[i] = b
	}
	hash := binary.BigEndian.Uint64(bytes)
	return uint(hash) & (numBuckets - 1)
}
