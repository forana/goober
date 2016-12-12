package flags

// Z is set if an operation was zero, whatever that means in context.
const Z = uint8(0x01) << 7

// N is set if the operation was subtraction.
const N = uint8(0x01) << 6

// H is set if the bottom nibble of an operation required a carry.
const H = uint8(0x01) << 5

// C is set if the bottom byte of an operation required a carry.
const C = uint8(0x01) << 4
