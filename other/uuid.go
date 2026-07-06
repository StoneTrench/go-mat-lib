package other

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

type UUID [4]uint32

func NewUUID() (UUID, error) {
	var b [16]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return UUID{}, fmt.Errorf("UUID failed to initialize random, %w", err)
	}

	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80

	var uuid UUID
	uuid[0] = binary.BigEndian.Uint32(b[0:4])
	uuid[1] = binary.BigEndian.Uint32(b[4:8])
	uuid[2] = binary.BigEndian.Uint32(b[8:12])
	uuid[3] = binary.BigEndian.Uint32(b[12:16])

	return uuid, nil
}
func (u UUID) ToFixedSlice() [4]int32 {
	return [4]int32{
		int32(u[0]),
		int32(u[1]),
		int32(u[2]),
		int32(u[3]),
	}
}
func (u UUID) ToString() string {
	var b [16]byte
	binary.BigEndian.PutUint32(b[0:4], u[0])
	binary.BigEndian.PutUint32(b[4:8], u[1])
	binary.BigEndian.PutUint32(b[8:12], u[2])
	binary.BigEndian.PutUint32(b[12:16], u[3])
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func NewUUIDPanic() UUID {
	res, err := NewUUID()
	if err != nil {
		panic(err)
	}
	return res
}
