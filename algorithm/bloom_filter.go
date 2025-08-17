package algorithm

import (
	"hash"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	bitset    []bool
	k         uint
	n         uint
	m         uint
	hashFuncs []hash.Hash64
}

func New(size uint) *BloomFilter {
	return &BloomFilter{
		bitset:    make([]bool, size),
		k:         3, // 3 hash funcs
		m:         size,
		n:         uint(0),
		hashFuncs: []hash.Hash64{murmur3.New64(), fnv.New64(), fnv.New64a()},
	}
}

func (bf *BloomFilter) Add(item []byte) {
	hashes := bf.hashValues(item)
	i := uint(0)
	for {
		if i >= bf.k {
			break
		}

		pos := uint(hashes[i]) % bf.m
		bf.bitset[uint(pos)] = true
		i += 1
	}

	bf.n += 1
}

func (bf *BloomFilter) Contains(item []byte) (exists bool) {
	hashes := bf.hashValues(item)
	i := uint(0)
	exists = true

	for {
		if i >= bf.k {
			break
		}

		position := uint(hashes[i]) % bf.m
		if !bf.bitset[uint(position)] {
			exists = false
			break
		}
		i += 1
	}
	return
}

func (bf *BloomFilter) hashValues(item []byte) []uint64 {
	var result []uint64
	for _, hashFunc := range bf.hashFuncs {
		hashFunc.Write(item)
		result = append(result, hashFunc.Sum64())
		hashFunc.Reset()
	}
	return result
}

// func main() {}
