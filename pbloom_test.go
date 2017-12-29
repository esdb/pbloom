package pbloom

import (
	"testing"
	"github.com/esdb/biter"
	"github.com/stretchr/testify/require"
)

func Test_one_location(t *testing.T) {
	should := require.New(t)
	strategy := NewHashingStrategy(HasherFnv, 256, 1)
	pbf := strategy.New()
	slot17 := biter.SetBits[17]
	strategy.Put(pbf, slot17, []byte("hello"))
	result := strategy.Find(pbf, []byte("hello"))
	should.NotEqual(biter.Bits(0), result)
	should.Equal(17, result.ScanForward()())
	result = strategy.Find(pbf, []byte("world2"))
	should.Equal(biter.Bits(0), result)
}

func Test_seven_locations(t *testing.T) {
	should := require.New(t)
	strategy := NewHashingStrategy(HasherFnv,256, 7)
	pbf := strategy.New()
	slot17 := biter.SetBits[17]
	strategy.Put(pbf, slot17, []byte("hello"))
	result := strategy.Find(pbf, []byte("hello"))
	should.NotEqual(0, result)
	should.Equal(17, result.ScanForward()())
}

func Test_two_slots(t *testing.T) {
	should := require.New(t)
	strategy := NewHashingStrategy(HasherFnv, 256, 1)
	pbf := strategy.New()
	slot17 := biter.SetBits[17]
	strategy.Put(pbf, slot17, []byte("hello"))
	slot18 := biter.SetBits[18]
	strategy.Put(pbf, slot18, []byte("hello"))
	result := strategy.Find(pbf, []byte("hello"))
	should.NotEqual(biter.Bits(0), result)
	iter := result.ScanForward()
	should.Equal(17, iter())
	should.Equal(18, iter())
	should.Equal(SlotNotFound, iter())
}

func Test_pre_hashing(t *testing.T) {
	should := require.New(t)
	strategy := NewHashingStrategy(HasherFnv, 256, 1)
	pbf := strategy.New()
	element := strategy.Hash([]byte("hello"))
	slot18 := biter.SetBits[18]
	pbf.Put(slot18, element)
	result := pbf.Find(element)
	should.NotEqual(biter.Bits(0), result)
	iter := result.ScanForward()
	should.Equal(18, iter())
	should.Equal(64, iter())
}
