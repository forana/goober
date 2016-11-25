package commands

import (
	"github.com/forana/goober/debug"
	"github.com/forana/goober/state"
)

// returns a pointer and the added cycle count to access it
func shortcodeToPointer(s *state.State, shortcode uint8) (*uint8, int) {
	switch shortcode {
	case 0x0:
		return &s.RAM.B, 0
	case 0x1:
		return &s.RAM.C, 0
	case 0x2:
		return &s.RAM.D, 0
	case 0x3:
		return &s.RAM.E, 0
	case 0x4:
		return &s.RAM.H, 0
	case 0x5:
		return &s.RAM.L, 0
	case 0x6:
		return s.RAM.HL(), 4
	case 0x7:
		return &s.RAM.A, 0
	default:
		debug.LogShortcodeError(shortcode)
		return debug.Slop, 0
	}
}

// Load00 loads an immediate 8-bit value into an arbitrary register.
// LD[D][I] - 0b00rrrmm0 [n]
func Load00(s *state.State, opcode uint8) int {
	switch opcode & 0x07 {
	case 0x00:
		addr := (uint16(s.Read()) << 8) | uint16(s.Read())
		s.RAM.AddressSpace[addr] = uint8(s.RAM.SP >> 8)
		s.RAM.AddressSpace[addr+1] = uint8(s.RAM.SP)
		return 20
	case 0x01:
		if opcode>>4 == 0x06 {
			s.RAM.SP = (uint16(s.Read()) << 8) | uint16(s.Read())
		} else {
			switch opcode >> 4 {
			case 0x00:
				s.RAM.B = s.Read()
				s.RAM.C = s.Read()
			case 0x02:
				s.RAM.D = s.Read()
				s.RAM.E = s.Read()
			case 0x04:
				s.RAM.H = s.Read()
				s.RAM.L = s.Read()
			default:
				debug.LogInstructionProblem("LD-16", opcode, "")
			}
		}
		return 12
	case 0x02:
		var r1 *uint8
		var r2 *uint8
		switch (opcode >> 3) & 0x07 {
		case 0x00:
			r1 = s.RAM.BC()
			r2 = &s.RAM.A
		case 0x01:
			r1 = &s.RAM.A
			r2 = s.RAM.BC()
		case 0x02:
			r1 = s.RAM.DE()
			r2 = &s.RAM.A
		case 0x03:
			r1 = &s.RAM.A
			r2 = s.RAM.DE()
		case 0x04:
			r1 = s.RAM.HL()
			r2 = &s.RAM.A
			s.RAM.L++
			if s.RAM.L == 0x00 {
				s.RAM.H++
			}
		case 0x05:
			r1 = &s.RAM.A
			r2 = s.RAM.HL()
			s.RAM.L++
			if s.RAM.L == 0x00 {
				s.RAM.H++
			}
		case 0x06:
			r1 = s.RAM.HL()
			r2 = &s.RAM.A
			s.RAM.L--
			if s.RAM.L == 0xFF {
				s.RAM.H--
			}
		case 0x07:
			r1 = &s.RAM.A
			r2 = s.RAM.HL()
			s.RAM.L--
			if s.RAM.L == 0xFF {
				s.RAM.H--
			}
		default:
			debug.LogInstructionProblem("LD[DI]", opcode, "")
			r1 = debug.Slop
			r2 = debug.Slop
		}
		*r1 = *r2
		return 8
	case 0x06:
		r, c := shortcodeToPointer(s, (opcode>>3)&0x07)
		*r = s.Read()
		return 4 + c
	default:
		debug.LogInstructionProblem("LD?", opcode, "")
		return 8
	}
}

// Load01 copies the value at one location into another.
// LD - 0b01rrrRRR
func Load01(s *state.State, opcode uint8) int {
	r1, c1 := shortcodeToPointer(s, (opcode>>3)&0x07)
	r2, c2 := shortcodeToPointer(s, opcode&0x07)
	*r1 = *r2
	return 4 + c1 + c2
}

// Load11 loads to/from A to/from an immediate (lsb first) address (or stack pointer).
// LD - 0b111mm010 [lsb msb]
func Load11(s *state.State, opcode uint8) int {
	var r1 *uint8
	var r2 *uint8
	var c = 8
	switch (opcode >> 3) & 0xFC {
	case 0x00:
		r1 = &s.RAM.AddressSpace[0xFF00|uint16(s.RAM.C)]
		r2 = &s.RAM.A
	case 0x01:
		r1 = s.RAM.LSFirst(s.Read(), s.Read())
		r2 = &s.RAM.A
		c += 8
	case 0x02:
		r1 = &s.RAM.A
		r2 = &s.RAM.AddressSpace[0xFF00|uint16(s.RAM.C)]
	case 0x03:
		r1 = &s.RAM.A
		r2 = s.RAM.LSFirst(s.Read(), s.Read())
		c += 8
	case 0x07:
		if opcode&0x01 == 0x01 {
			r1 = &s.RAM.AddressSpace[s.RAM.SP]
			r2 = s.RAM.HL()
		} else {
			return loadHLN(s, opcode)
		}
	default:
		debug.LogInstructionProblem("LoadAImmediate", opcode, "")
		r1 = debug.Slop
		r2 = debug.Slop
	}
	*r1 = *r2
	return c
}

func loadHLN(s *state.State, opcode uint8) int {
	s.RAM.F = 0x00
	n := int16(s.Read())
	sp := s.RAM.SP
	nsp := uint16(int16(sp) + n)
	if (nsp < sp && n > 0) || (nsp > sp && n < 0) {
		s.RAM.F |= flagCarry
	}

	return 12
}
