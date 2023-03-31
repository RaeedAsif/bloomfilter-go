package storage

import (
	"hash"
	"hash/fnv"
	"math"
	"sync"

	"github.com/RaeedAsif/bloomfilter-go/utils"

	"github.com/spaolacci/murmur3"
)

// ref : https://codeburst.io/lets-implement-a-bloom-filter-in-go-b2da8a4b849f

var (
	bf            *BloomFilter
	ExpectedUsers = 100000
	FalsePositive = 0.1 // 10%
)

// BloomFilter struct
type BloomFilter struct {
	bitset    []bool        // The bloom-filter bitset
	m         uint          // Size of the bloom filter
	hashfuncs []hash.Hash64 // The hash functions
	mu        sync.Mutex
}

var DefaultHashFunctions = []hash.Hash64{murmur3.New64(), fnv.New64(), fnv.New64a()}

// Returns a new BloomFilter object
func InitBloomFilter() {
	size := getBitSize()
	bf = &BloomFilter{
		bitset:    make([]bool, size),
		m:         uint(size),
		hashfuncs: DefaultHashFunctions,
	}
}

// Returns a new BloomFilter object
func GetInstance() *BloomFilter {
	return bf
}

// Sets the item into the bloom filter set by hashing in over the hash functions
func (bf *BloomFilter) Set(username string) {
	item := []byte(username)
	for _, v := range bf.get(item) {
		position := uint(v) % bf.m
		bf.bitset[position] = true
	}
}

// Check if the item into the bloom filter is set by hashing in over the hash functions
func (bf *BloomFilter) Check(username string) bool {
	item := []byte(username)
	for _, v := range bf.get(item) {
		position := uint(v) % bf.m
		if !bf.bitset[uint(position)] {
			return false
		}
	}
	return true
}

// Get all the hash values by hashing in over the hash functions
func (bf *BloomFilter) get(item []byte) []uint64 {
	var result []uint64

	for _, hashFunc := range bf.hashfuncs {
		hashFunc.Write(item)
		result = append(result, hashFunc.Sum64())
		hashFunc.Reset()
	}

	return result
}

// Returns bit size
/*
Legend:
n: how many items you expect to have in your filter (e.g. 216,553)
p: your acceptable false positive rate {0..1} (e.g. 0.01 → 1%)
m: the number of bits needed in the bloom filter
k: the number of hash functions we should apply

Formulas:
m = -n*ln(p) / (ln(2)^2) the number of bits
k = m/n * ln(2) the number of hash functions
*/
func getBitSize() int {
	return int((float64(len(DefaultHashFunctions)) * math.Pow(math.Ln2, 2)) / utils.Ln(FalsePositive))
}
