package state

import (
	"github.com/forana/goober/flags"
	"github.com/forana/goober/memory"
	"github.com/forana/goober/rom"
)

// State represents the entire running state of the system, as it knows itself
type State struct {
	RAM  *memory.RAM
	VRAM *memory.VRAM
	ROM  *rom.ROM
}

// Read reads in a byte an increments the program counter.
func (s *State) Read() uint8 {
	b := s.ROM.Program[s.RAM.PC]
	s.RAM.PC++
	return b
}

// Add8 adds n to v, setting any flags appropriately.
func (s *State) Add8(v uint8, nums ...uint8) uint8 {
	n := uint16(0)
	for i := range nums {
		n += uint16(i)
	}
	s.RAM.F = 0x00
	if uint16(v&0xF)+n >= 0x10 {
		s.RAM.F |= flags.H
	}
	if uint(v)+uint(n) > 0xFF {
		s.RAM.F |= flags.C
	}
	r := uint8((uint16(v) + n) & 0xFF)
	if r == 0 {
		s.RAM.F |= flags.Z
	}
	return r
}

// Add16Signed adds signed n to unsigned v, setting carry flags appropriately.
func (s *State) Add16Signed(v uint16, n int8) uint16 {
	s.RAM.F = 0x00
	if int16(v&0xF)+int16(n) >= 0x10 {
		s.RAM.F |= flags.H
	}
	if int(v)+int(n) > 0xFF {
		s.RAM.F |= flags.C
	}
	return uint16(uint(int(v)+int(n)) & 0xFFFF)
}

// FlagSet checks if the given flag is set.
func (s *State) FlagSet(flag uint8) bool {
	return s.RAM.F&flag != 0
}
