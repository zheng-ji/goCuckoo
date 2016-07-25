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

func genSignature(data []byte) Signature {
	hashInstance := fnv.New64()
	hashInstance.Reset()
	hashInstance.Write(data)
	hash := hashInstance.Sum(nil)
	sign := Signature{}
	for i := 0; i < SignatureSize; i++ {
		sign[i] = hash[i]
	}
	if sign == Empty {
		sign[0] ^= 1
	}
	return sign
}

func genFirstIndex(sign Signature, numBuckets uint) uint {
	bytes := make([]byte, 64, 64)
	for i, b := range sign {
		bytes[i] = b
	}
	hash := binary.LittleEndian.Uint64(bytes)
	return uint(hash) & (numBuckets - 1)
}

func genBackupIndex(sign Signature, numBuckets uint) uint {
	bytes := make([]byte, 64, 64)
	for i, b := range sign {
		bytes[i] = b
	}
	hash := binary.BigEndian.Uint64(bytes)
	return uint(hash) & (numBuckets - 1)
}
