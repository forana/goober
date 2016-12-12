package memory

const vramAddressMax = 65536

// VRAM inside the console
type VRAM struct {
	Memory [vramAddressMax]uint8
}

// InitVRAM creates a new VRAM model
func InitVRAM() *VRAM {
	// TODO init memory properly
	return &VRAM{}
}
