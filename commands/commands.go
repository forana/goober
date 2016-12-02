package commands

import (
	"github.com/forana/goober/memory"
	"github.com/forana/goober/state"
)

const flagZero = uint8(0x01) << 7
const flagNSubtracted = uint8(0x01) << 6
const flagHalfCarry = uint8(0x01) << 5
const flagCarry = uint8(0x01) << 4

// Command is a thing.
type Command func(*state.State) uint

// Registry maps opcodes to commands.
func Registry() map[uint8]Command {
	reg := make(map[uint8]Command)

	reg[0x06] = load8BitsToRegister(memory.B)
	reg[0x0E] = load8BitsToRegister(memory.C)
	reg[0x16] = load8BitsToRegister(memory.D)
	reg[0x1E] = load8BitsToRegister(memory.E)
	reg[0x26] = load8BitsToRegister(memory.H)
	reg[0x2E] = load8BitsToRegister(memory.L)

	reg[0x7F] = loadCopy(memory.A, memory.A)
	reg[0x78] = loadCopy(memory.A, memory.B)
	reg[0x79] = loadCopy(memory.A, memory.C)
	reg[0x7A] = loadCopy(memory.A, memory.D)
	reg[0x7B] = loadCopy(memory.A, memory.E)
	reg[0x7C] = loadCopy(memory.A, memory.H)
	reg[0x7D] = loadCopy(memory.A, memory.L)
	reg[0x7E] = loadHLLocationToRegister(memory.A)
	reg[0x40] = loadCopy(memory.B, memory.B)
	reg[0x41] = loadCopy(memory.B, memory.C)
	reg[0x42] = loadCopy(memory.B, memory.D)
	reg[0x43] = loadCopy(memory.B, memory.E)
	reg[0x44] = loadCopy(memory.B, memory.H)
	reg[0x45] = loadCopy(memory.B, memory.L)
	reg[0x46] = loadHLLocationToRegister(memory.B)
	reg[0x48] = loadCopy(memory.C, memory.B)
	reg[0x49] = loadCopy(memory.C, memory.C)
	reg[0x4A] = loadCopy(memory.C, memory.D)
	reg[0x4B] = loadCopy(memory.C, memory.E)
	reg[0x4C] = loadCopy(memory.C, memory.H)
	reg[0x4D] = loadCopy(memory.C, memory.L)
	reg[0x4E] = loadHLLocationToRegister(memory.C)
	reg[0x50] = loadCopy(memory.D, memory.B)
	reg[0x51] = loadCopy(memory.D, memory.C)
	reg[0x52] = loadCopy(memory.D, memory.D)
	reg[0x53] = loadCopy(memory.D, memory.E)
	reg[0x54] = loadCopy(memory.D, memory.H)
	reg[0x55] = loadCopy(memory.D, memory.L)
	reg[0x56] = loadHLLocationToRegister(memory.D)
	reg[0x58] = loadCopy(memory.E, memory.B)
	reg[0x59] = loadCopy(memory.E, memory.C)
	reg[0x5A] = loadCopy(memory.E, memory.D)
	reg[0x5B] = loadCopy(memory.E, memory.E)
	reg[0x5C] = loadCopy(memory.E, memory.H)
	reg[0x5D] = loadCopy(memory.E, memory.L)
	reg[0x5E] = loadHLLocationToRegister(memory.E)
	reg[0x60] = loadCopy(memory.H, memory.B)
	reg[0x61] = loadCopy(memory.H, memory.C)
	reg[0x62] = loadCopy(memory.H, memory.D)
	reg[0x63] = loadCopy(memory.H, memory.E)
	reg[0x64] = loadCopy(memory.H, memory.H)
	reg[0x65] = loadCopy(memory.H, memory.L)
	reg[0x66] = loadHLLocationToRegister(memory.H)
	reg[0x68] = loadCopy(memory.L, memory.B)
	reg[0x69] = loadCopy(memory.L, memory.C)
	reg[0x6A] = loadCopy(memory.L, memory.D)
	reg[0x6B] = loadCopy(memory.L, memory.E)
	reg[0x6C] = loadCopy(memory.L, memory.H)
	reg[0x6D] = loadCopy(memory.L, memory.L)
	reg[0x6E] = loadHLLocationToRegister(memory.L)
	reg[0x70] = loadRegisterToCompositeLocation(memory.B, memory.H, memory.L)
	reg[0x71] = loadRegisterToCompositeLocation(memory.C, memory.H, memory.L)
	reg[0x72] = loadRegisterToCompositeLocation(memory.D, memory.H, memory.L)
	reg[0x73] = loadRegisterToCompositeLocation(memory.E, memory.H, memory.L)
	reg[0x74] = loadRegisterToCompositeLocation(memory.H, memory.H, memory.L)
	reg[0x75] = loadRegisterToCompositeLocation(memory.L, memory.H, memory.L)
	reg[0x36] = loadImmediateToHL

	return reg
}
