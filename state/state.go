package state

import (
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
