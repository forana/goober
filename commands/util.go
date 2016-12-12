package commands

import (
	"unsafe"
)

func make16(r1 uint8, r2 uint8) uint16 {
	return (uint16(r1) << 8) | uint16(r2)
}

func make16LSFirst(r1 uint8, r2 uint8) uint16 {
	return make16(r2, r1)
}

func signed(x uint8) int8 {
	return *(*int8)(unsafe.Pointer(&x))
}

func unsigned(x int8) uint8 {
	return *(*uint8)(unsafe.Pointer(&x))
}

func top(v uint16) uint8 {
	return uint8(v >> 8)
}

func bottom(v uint16) uint8 {
	return uint8(v)
}

func add8(a uint8, b int) uint8 {
	return uint8((int(a) + b) % 0x100)
}

func add16(a uint16, b int) uint16 {
	return uint16((int(a) + b) % 0x10000)
}

func cbool(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
