package models

import (
	"encoding/binary"
)

// Size of control bytes buffer
const (
	ControlBytesSize = 32
)

// ControlBytes ...
type ControlBytes []byte

// GetXY ...
func (b *ControlBytes) GetXY() (int, int) {
	if b == nil || len(*b) != ControlBytesSize {
		return 0, 0
	}

	x := binary.BigEndian.Uint32((*b)[10:14])
	y := binary.BigEndian.Uint32((*b)[14:18])
	return int(x), int(y)
}

// ApplyOffset ...
func (b *ControlBytes) ApplyOffset(x, y int) {
	if b == nil || len(*b) != ControlBytesSize || (x == 0 && y == 0) {
		return
	}

	var oldX = binary.BigEndian.Uint32((*b)[10:14])
	var oldY = binary.BigEndian.Uint32((*b)[14:18])

	var xWithOffset = int(oldX) + x
	if xWithOffset < 0 {
		xWithOffset = 0
	}

	var yWithOffset = int(oldY) + y
	if yWithOffset < 0 {
		yWithOffset = 0
	}

	binary.BigEndian.PutUint32((*b)[10:14], uint32(xWithOffset))
	binary.BigEndian.PutUint32((*b)[14:18], uint32(yWithOffset))
}
