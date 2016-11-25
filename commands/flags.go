package commands

// flagZero is set when a math operation is equal.
const flagZero = uint8(0x01) << 7

// flagNSubtracted is set when subtraction was performed in the last operation.
const flagNSubtracted = uint8(0x01) << 6

// flagHalfCarry is set when a carry happened from the lower nibble during a math operation.
const flagHalfCarry = uint8(0x01) << 5

// flagCarry is set when a carry happened in the last operation or if A was the smaller value when compared.
const flagCarry = uint8(0x01) << 4
