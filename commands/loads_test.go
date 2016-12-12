package commands

import (
	"testing"

	"github.com/forana/goober/assert"
	"github.com/forana/goober/memory"
	"github.com/forana/goober/rom"
	"github.com/forana/goober/state"
)

func newState() *state.State {
	return &state.State{
		RAM:  memory.InitRAM(),
		VRAM: memory.InitVRAM(),
		ROM:  &rom.ROM{Program: make([]uint8, 65536)},
	}
}

func TestLoadImmediateToRegister(t *testing.T) {
	s := newState()
	value := uint8(0xAA)
	s.RAM.B = 0x00
	s.ROM.Program[s.RAM.PC] = value
	assert.NotEqual(t, value, s.RAM.B)
	loadImmediateToRegister(memory.B)(s)
	assert.Equal(t, value, s.RAM.B)
}

func TestLoadCopy(t *testing.T) {
	s := newState()
	s.RAM.B = 0x11
	s.RAM.C = 0x22
	assert.NotEqual(t, s.RAM.B, s.RAM.C)
	loadCopy(memory.B, memory.C)(s)
	assert.Equal(t, uint8(0x22), s.RAM.B)
}

func TestLoadRegisterToHLLocation(t *testing.T) {
	s := newState()
	s.RAM.H = 0x12
	s.RAM.L = 0x34
	s.RAM.Memory[0x1234] = 0xAA
	s.RAM.A = 0xBB
	assert.NotEqual(t, uint8(0xBB), s.RAM.Memory[0x1234])
	loadRegisterToHLLocation(memory.A)(s)
	assert.Equal(t, uint8(0xBB), s.RAM.Memory[0x1234])
}

func TestLoadCompositeLocationToRegister(t *testing.T) {
	s := newState()
	s.RAM.A = 0x11
	s.RAM.B = 0x43
	s.RAM.C = 0x21
	s.RAM.Memory[0x4321] = 0x22
	assert.NotEqual(t, uint8(0x22), s.RAM.A)
	loadCompositeLocationToRegister(memory.A, memory.B, memory.C)(s)
	assert.Equal(t, uint8(0x22), s.RAM.A)
}

func TestLoadImmediateToHLLocation(t *testing.T) {
	s := newState()
	s.RAM.H = 0x12
	s.RAM.L = 0x34
	s.RAM.Memory[0x1234] = 0x11
	s.ROM.Program[s.RAM.PC] = 0x22
	assert.NotEqual(t, uint8(0x22), s.RAM.Memory[0x1234])
	loadImmediateToHLLocation(s)
	assert.Equal(t, uint8(0x22), s.RAM.Memory[0x1234])
}

func TestLoadImmediateLocationToA(t *testing.T) {
	s := newState()
	s.RAM.A = 0xAA
	s.ROM.Program[s.RAM.PC] = 0x34
	s.ROM.Program[s.RAM.PC+1] = 0x12
	s.RAM.Memory[0x1234] = 0xBB
	assert.NotEqual(t, uint8(0xBB), s.RAM.A)
	loadImmediateLocationToA(s)
	assert.Equal(t, uint8(0xBB), s.RAM.A)
}
