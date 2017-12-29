package pbloom

import (
	"testing"
	"github.com/esdb/biter"
	"github.com/stretchr/testify/require"
)

func Test_one_location(t *testing.T) {
	should := require.New(t)
	pbf := New(256, 1)
	slot17 := biter.SetBits[17]
	pbf.Put(HasherFnv, slot17, []byte("hello"))
	result := pbf.Find(HasherFnv, []byte("hello"))
	should.NotEqual(biter.Bits(0), result)
	should.Equal(17, result.ScanForward()())
	result = pbf.Find(HasherFnv, []byte("world2"))
	should.Equal(biter.Bits(0), result)
}

func Test_seven_locations(t *testing.T) {
	should := require.New(t)
	pbf := New(256, 7)
	slot17 := biter.SetBits[17]
	pbf.Put(HasherFnv, slot17, []byte("hello"))
	result := pbf.Find(HasherFnv, []byte("hello"))
	should.NotEqual(0, result)
	should.Equal(17, result.ScanForward()())
}

func Test_two_slots(t *testing.T) {
	should := require.New(t)
	pbf := New(256, 1)
	slot17 := biter.SetBits[17]
	pbf.Put(HasherFnv, slot17, []byte("hello"))
	slot18 := biter.SetBits[18]
	pbf.Put(HasherFnv, slot18, []byte("hello"))
	result := pbf.Find(HasherFnv, []byte("hello"))
	should.NotEqual(biter.Bits(0), result)
	iter := result.ScanForward()
	should.Equal(17, iter())
	should.Equal(18, iter())
	should.Equal(SlotNotFound, iter())
}
