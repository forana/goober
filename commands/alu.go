package commands

import (
	"github.com/forana/goober/flags"
	"github.com/forana/goober/memory"
	"github.com/forana/goober/state"
)

// add a register's value to A
func addRegisterToA(r memory.Register) Command {
	return func(s *state.State) uint {
		s.RAM.A = s.Add8(s.RAM.A, *s.RAM.Register(r))
		return 4
	}
}

// add the value at HL to A
func addHLToA(s *state.State) uint {
	s.RAM.A = s.Add8(s.RAM.A, s.RAM.Memory[make16(s.RAM.H, s.RAM.L)])
	return 8
}

// add immediate value to A
func addImmediateToA(s *state.State) uint {
	s.RAM.A = s.Add8(s.RAM.A, s.Read())
	return 8
}

// add carry + a register's value to A
func addCarryRegisterToA(r memory.Register) Command {
	return func(s *state.State) uint {
		s.RAM.A = s.Add8(s.RAM.A, *s.RAM.Register(r), cbool(s.FlagSet(flags.C)))
		return 4
	}
}

// add carry + the value at HL to A
func addCarryHLToA(s *state.State) uint {
	s.RAM.A = s.Add8(s.RAM.A, s.RAM.Memory[make16(s.RAM.H, s.RAM.L)], cbool(s.FlagSet(flags.C)))
	return 8
}

// add carry + immediate value to A
func addCarryImmediateToA(s *state.State) uint {
	s.RAM.A = s.Add8(s.RAM.A, s.Read(), cbool(s.FlagSet(flags.C)))
	return 8
}
