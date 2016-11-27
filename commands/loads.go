package commands

import (
	"github.com/forana/goober/memory"
	"github.com/forana/goober/state"
)

// 8-bit load
func load8BitsToRegister(r memory.Register) Command {
	return func(s *state.State) uint {
		register := s.RAM.Register(r)
		*register = s.Read()
		return 8
	}
}

// copy register to register
func loadCopy(r1 memory.Register, r2 memory.Register) Command {
	return func(s *state.State) uint {
		*(s.RAM.Register(r1)) = *(s.RAM.Register(r2))
		return 4
	}
}

// copy register valye to memory location stored in HL
func loadRegisterToHLLocation(r memory.Register) Command {
	return func(s *state.State) uint {
		s.RAM.AddressSpace[s.RAM.Get16(&s.RAM.H, &s.RAM.L)] = *(s.RAM.Register(r))
		return 8
	}
}

// copy from memory location stored in HL to register
func loadHLLocationToRegister(r memory.Register) Command {
	return func(s *state.State) uint {
		*(s.RAM.Register(r)) = s.RAM.AddressSpace[s.RAM.Get16(&s.RAM.H, &s.RAM.L)]
		return 8
	}
}

// load immediate 8 bits into memory location stored in HL {
func loadImmediateToHL(s *state.State) uint {
	s.RAM.AddressSpace[s.RAM.Get16(&s.RAM.H, &s.RAM.L)] = s.Read()
	return 12
}
