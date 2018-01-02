package pbloom

import "unsafe"

// fnv is fast for small data
// xxhash is fast for larger data

const (
	offset64        = 14695981039346656037
	prime64         = 1099511628211
)

// copied stdlib hash/fnv
func HasherFnv(element []byte) (s HashedElement) {
	hash := uint64(offset64)
	for _, c := range element {
		hash ^= uint64(c)
		hash *= prime64
	}
	return *(*[2]uint32)(unsafe.Pointer(&hash))
}
