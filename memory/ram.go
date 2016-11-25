package memory

import (
	"math/rand"
)

const ramAddressMax = 65536

// RAM inside the console
type RAM struct {
	AddressSpace [ramAddressMax]uint8

	// general registers for programming - not part of ram??
	AF uint8
	F  uint8
	B  uint8
	C  uint8
	D  uint8
	E  uint8
	H  uint8
	L  uint8

	SP uint16 // stack pointer

	P1   *uint8 // joypad info
	SB   *uint8 // serial transfer data
	SC   *uint8 // S10 control
	DIV  *uint8 // incremented 16384 times per second, any write sets it to 0x00
	TIMA *uint8 // timer counter, generates interrupt on overflow
	TMA  *uint8 // timer modulo, set on TIMA overflow
	TAC  *uint8 // timer control
	IF   *uint8 // interrupt flag
	NR10 *uint8 // sound mode 1
	NR11 *uint8 // sound mode 1
	NR12 *uint8 // sound mode 1
	NR13 *uint8 // sound mode 1
	NR14 *uint8 // sound mode 1
	NR21 *uint8 // sound mode 2
	NR22 *uint8 // sound mode 2
	NR23 *uint8 // sound mode 2
	NR24 *uint8 // sound mode 2
	NR30 *uint8 // sound mode 3
	NR31 *uint8 // sound mode 3
	NR32 *uint8 // sound mode 3
	NR33 *uint8 // sound mode 3
	NR34 *uint8 // sound mode 3
	NR41 *uint8 // sound mode 4
	NR42 *uint8 // sound mode 4
	NR43 *uint8 // sound mode 4
	NR44 *uint8 // sound mode 4
	NR50 *uint8 // channel control / volume
	NR51 *uint8 // sound output terminal selection
	NR52 *uint8 // sound on/off
	LCDC *uint8 // LCD control
	STAT *uint8 // LCD status, set by operation
	SCY  *uint8 // scroll Y
	SCX  *uint8 // scroll X
	LY   *uint8 // LCDC y-coordinate
	LYC  *uint8 // LY compare (coincidence)
	DMA  *uint8 // DMA transfer and start address
	BGP  *uint8 // bg/window palette data
	OBP0 *uint8 // object palette 0 data
	OBP1 *uint8 // object palette 1 data
	WY   *uint8 // window Y pos
	WX   *uint8 // window X pos
	IE   *uint8 // interrupt enable
}

// InitRAM creates a new RAM model, pretending to be a GBC
func InitRAM() *RAM {
	// init registers + stack pointer
	ram := &RAM{
		AF: 0x11,
		F:  0xB0,
		B:  0x00,
		C:  0x13,
		D:  0x00,
		E:  0xD8,
		H:  0x01,
		L:  0x4D,
		SP: 0xFFFE,
	}
	// fill ram with garbage
	for a := 0; a < ramAddressMax; a++ {
		ram.AddressSpace[a] = uint8(rand.Intn(256))
	}
	// init aliases to special places
	ram.P1 = &ram.AddressSpace[0xFF00]
	ram.SB = &ram.AddressSpace[0xFF01]
	ram.SC = &ram.AddressSpace[0xFF02]
	ram.DIV = &ram.AddressSpace[0xFF04]
	ram.TIMA = &ram.AddressSpace[0xFF05]
	ram.TMA = &ram.AddressSpace[0xFF06]
	ram.TAC = &ram.AddressSpace[0xFF07]
	ram.IF = &ram.AddressSpace[0xFF0F]
	ram.NR10 = &ram.AddressSpace[0xFF10]
	ram.NR11 = &ram.AddressSpace[0xFF11]
	ram.NR12 = &ram.AddressSpace[0xFF12]
	ram.NR13 = &ram.AddressSpace[0xFF13]
	ram.NR14 = &ram.AddressSpace[0xFF14]
	ram.NR21 = &ram.AddressSpace[0xFF16]
	ram.NR22 = &ram.AddressSpace[0xFF17]
	ram.NR23 = &ram.AddressSpace[0xFF18]
	ram.NR24 = &ram.AddressSpace[0xFF19]
	ram.NR30 = &ram.AddressSpace[0xFF1A]
	ram.NR31 = &ram.AddressSpace[0xFF1B]
	ram.NR32 = &ram.AddressSpace[0xFF1C]
	ram.NR33 = &ram.AddressSpace[0xFF1D]
	ram.NR34 = &ram.AddressSpace[0xFF1E]
	ram.NR41 = &ram.AddressSpace[0xFF20]
	ram.NR42 = &ram.AddressSpace[0xFF21]
	ram.NR43 = &ram.AddressSpace[0xFF22]
	ram.NR44 = &ram.AddressSpace[0xFF23]
	ram.NR50 = &ram.AddressSpace[0xFF24]
	ram.NR51 = &ram.AddressSpace[0xFF25]
	ram.NR52 = &ram.AddressSpace[0xFF26]
	// 0xFF30 - 0xFF3F are wave pattern ram
	ram.LCDC = &ram.AddressSpace[0xFF40]
	ram.STAT = &ram.AddressSpace[0xFF41]
	ram.SCY = &ram.AddressSpace[0xFF42]
	ram.SCX = &ram.AddressSpace[0xFF43]
	ram.LY = &ram.AddressSpace[0xFF44]
	ram.LYC = &ram.AddressSpace[0xFF45]
	ram.DMA = &ram.AddressSpace[0xFF46]
	ram.BGP = &ram.AddressSpace[0xFF47]
	ram.OBP0 = &ram.AddressSpace[0xFF48]
	ram.OBP1 = &ram.AddressSpace[0xFF49]
	ram.WY = &ram.AddressSpace[0xFF4A]
	ram.WX = &ram.AddressSpace[0xFF4B]
	ram.IE = &ram.AddressSpace[0xFFFF]
	// init known special places to known hardware defaults
	*ram.TIMA = 0x00
	*ram.TMA = 0x00
	*ram.TAC = 0x00
	*ram.NR10 = 0x80
	*ram.NR11 = 0xBF
	*ram.NR12 = 0xF3
	*ram.NR14 = 0xBF
	*ram.NR21 = 0x3F
	*ram.NR22 = 0x00
	*ram.NR23 = 0x00 // ?
	*ram.NR24 = 0xBF
	*ram.NR30 = 0x7F
	*ram.NR31 = 0xFF
	*ram.NR32 = 0x9F
	*ram.NR33 = 0xBF
	*ram.NR34 = 0x00 // ?
	*ram.NR41 = 0xFF
	*ram.NR42 = 0x00
	*ram.NR43 = 0x00
	*ram.NR44 = 0xBF
	*ram.NR50 = 0x77
	*ram.NR51 = 0xF3
	*ram.NR52 = 0xF0
	*ram.LCDC = 0x91
	*ram.SCY = 0x00
	*ram.SCX = 0x00
	*ram.LY = 0x00 // ?
	*ram.LYC = 0x00
	*ram.DMA = 0x00 // ?
	*ram.BGP = 0xFC
	*ram.OBP0 = 0xFF
	*ram.OBP1 = 0xFF
	*ram.WY = 0x00
	*ram.WX = 0x00
	*ram.IE = 0x00

	return ram
}
