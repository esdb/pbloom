package pbloom

import (
	"math"
	"unsafe"
)

// element => hash1,hash2 => k locations
type Hasher func(element []byte) HashedElement

type HashedElement [2]uint64

type BloomElement []uint64

func (hashedElement HashedElement) Hash(locationsPerElement uint64, locationsCount uint64) BloomElement {
	kHashedElement := make(BloomElement, locationsPerElement)
	combinedHash := hashedElement[0]
	for i := uint64(0); i < locationsPerElement; i++ {
		kHashedElement[i] = (combinedHash & math.MaxUint64) % uint64(locationsCount)
		combinedHash += hashedElement[1]
	}
	return kHashedElement
}

func (hashedElement HashedElement) DownCastToUint32() uint32 {
	return *(*uint32)(unsafe.Pointer(&hashedElement[0]))
}
