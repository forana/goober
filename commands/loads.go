package commands

import (
	"github.com/forana/goober/memory"
	"github.com/forana/goober/state"
)

// 8-bit immediate load
func loadImmediateToRegister(r memory.Register) Command {
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

// copy register value to memory location stored in HL
func loadRegisterToCompositeLocation(r1 memory.Register, r2 memory.Register, r memory.Register) Command {
	return func(s *state.State) uint {
		s.RAM.Memory[make16(*s.RAM.Register(r1), *s.RAM.Register(r2))] = *(s.RAM.Register(r))
		return 8
	}
}

// copy from memory location stored in composite to register
func loadCompositeLocationToRegister(r memory.Register, r1 memory.Register, r2 memory.Register) Command {
	return func(s *state.State) uint {
		*(s.RAM.Register(r)) = s.RAM.Memory[make16(*s.RAM.Register(r1), *s.RAM.Register(r2))]
		return 8
	}
}

// load immediate 8 bits into memory location stored in HL
func loadImmediateToHLLocation(s *state.State) uint {
	s.RAM.Memory[make16(s.RAM.H, s.RAM.L)] = s.Read()
	return 12
}

// load location at immediate 16 bits (lsb first) into A
func loadImmediateLocationToA(s *state.State) uint {
	s.RAM.A = s.RAM.Memory[make16LSFirst(s.Read(), s.Read())]
	return 16
}

// load A into location at immediate 16 bits (lsb first)
func loadAToImmediateLocation(s *state.State) uint {
	s.RAM.Memory[make16LSFirst(s.Read(), s.Read())] = s.RAM.A
	return 16
}

// load from $FF00+C into A
func loadOffsetCIntoA(s *state.State) uint {
	s.RAM.A = s.RAM.Memory[add16(0xFF00, int(s.RAM.C))]
	return 8
}

// load from A into $FF00+C
func loadAIntoOffsetC(s *state.State) uint {
	s.RAM.Memory[add16(0xFF00, int(s.RAM.C))] = s.RAM.A
	return 8
}

// load from $FF00 + immediate 8-bit value into A
func loadOffsetImmediateIntoA(s *state.State) uint {
	s.RAM.A = s.RAM.Memory[add16(0xFF00, int(s.Read()))]
	return 12
}

// load from A into $FF00 + immediate 8-bit value
func loadAIntoOffsetImmediate(s *state.State) uint {
	s.RAM.Memory[add16(0xFF00, int(s.Read()))] = s.RAM.A
	return 12
}

// load value at location HL into A, then decrement HL
func loadDecrementHLToA(s *state.State) uint {
	hl := make16(s.RAM.H, s.RAM.L)
	s.RAM.A = s.RAM.Memory[hl]
	hl = add16(hl, -1)
	s.RAM.H = top(hl)
	s.RAM.L = bottom(hl)
	return 8
}

// load A to location HL, then decrement HL
func loadAToDecrementHL(s *state.State) uint {
	hl := make16(s.RAM.H, s.RAM.L)
	s.RAM.Memory[hl] = s.RAM.A
	hl = add16(hl, -1)
	s.RAM.H = top(hl)
	s.RAM.L = bottom(hl)
	return 8
}

// load value at location HL into A, then increment HL
func loadIncrementHLToA(s *state.State) uint {
	hl := make16(s.RAM.H, s.RAM.L)
	s.RAM.A = s.RAM.Memory[hl]
	hl = add16(hl, 1)
	s.RAM.H = top(hl)
	s.RAM.L = bottom(hl)
	return 8
}

// load A to location HL, then increment HL
func loadAToIncrementHL(s *state.State) uint {
	hl := make16(s.RAM.H, s.RAM.L)
	s.RAM.Memory[hl] = s.RAM.A
	hl = add16(hl, 1)
	s.RAM.H = top(hl)
	s.RAM.L = bottom(hl)
	return 8
}

// load immediate 16-bit value into composite location
func loadImmediateIntoComposite(r1 memory.Register, r2 memory.Register) Command {
	return func(s *state.State) uint {
		*s.RAM.Register(r1) = s.Read()
		*s.RAM.Register(r2) = s.Read()
		return 12
	}
}

// load immediate 16-bit value into SP
func loadImmediateIntoSP(s *state.State) uint {
	s.RAM.SP = make16(s.Read(), s.Read())
	return 12
}

// load value in HL into SP
func loadHLIntoSP(s *state.State) uint {
	s.RAM.SP = make16(s.RAM.H, s.RAM.L)
	return 8
}

// load SP + signed immediate 8-bit value into HL
func loadSPSignedImmediateOffsetIntoHL(s *state.State) uint {
	v := s.Add16Signed(s.RAM.SP, signed(s.Read()))
	s.RAM.H = top(v)
	s.RAM.L = bottom(v)
	return 12
}

// load SP to immediate location
func loadSPToImmediate(s *state.State) uint {
	l := make16(s.Read(), s.Read())
	s.RAM.Memory[l] = top(s.RAM.SP)
	s.RAM.Memory[add16(l, 1)] = bottom(s.RAM.SP)
	return 20
}

// push register pair onto stack
func push16(r1 memory.Register, r2 memory.Register) Command {
	return func(s *state.State) uint {
		s.RAM.Memory[s.RAM.SP] = *s.RAM.Register(r1)
		s.RAM.SP = add16(s.RAM.SP, -1)
		s.RAM.Memory[s.RAM.SP] = *s.RAM.Register(r2)
		s.RAM.SP = add16(s.RAM.SP, -1)
		return 16
	}
}

// pop from stack into register pair
func pop16(r1 memory.Register, r2 memory.Register) Command {
	return func(s *state.State) uint {
		*s.RAM.Register(r2) = s.RAM.Memory[s.RAM.SP]
		s.RAM.SP = add16(s.RAM.SP, 1)
		*s.RAM.Register(r1) = s.RAM.Memory[s.RAM.SP]
		s.RAM.SP = add16(s.RAM.SP, 1)
		return 12
	}
}
