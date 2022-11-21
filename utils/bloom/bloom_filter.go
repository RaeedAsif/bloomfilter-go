package bloom

import (
	"hash"
	"hash/fnv"
	"math"

	"github.com/spaolacci/murmur3"
)

var (
	bf            *BloomFilter
	ExpectedUsers = 100000
	FalsePositive = 0.1 // 10%
)

// BloomFilter
type BloomFilter struct {
	bitset    []bool        // The bloom-filter bitset
	m         uint          // Size of the bloom filter
	hashfuncs []hash.Hash64 // The hash functions
}

var DefaultHashFunctions = []hash.Hash64{murmur3.New64(), fnv.New64(), fnv.New64a()}

// Returns a new BloomFilter object
func Init() {
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
func getBitSize() int {
	return int((float64(len(DefaultHashFunctions)) * math.Pow(math.Ln2, 2)) / ln(FalsePositive))
}

// to calculate natural log
func f(x, a float64) float64 {

	return math.Exp(x) - a
}

func ln(n float64) float64 {
	var lo, hi, m float64

	if n <= 0 {
		return -1
	}

	if n == 1 {
		return 0
	}

	EPS := 0.00001
	lo = 0
	hi = n

	for math.Abs(lo-hi) >= EPS {
		m = float64((lo + hi) / 2.0)

		if f(m, n) < 0 {
			lo = m
		} else {
			hi = m
		}
	}

	return float64((lo + hi) / 2.0)
}
