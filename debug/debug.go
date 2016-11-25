package debug

import (
	"fmt"
)

// Slop is a spot to throw pointers to when things go wrong in parsing.
var Slop *uint8

// LogInstructionProblem logs a processing issue with an instruction.
// If this whole thing works, this will never be called.
func LogInstructionProblem(instruction string, opcode uint8, more string) {
	fmt.Printf("[inst] %s (%x) (%s)\n", instruction, opcode, more)
}

func LogShortcodeError(value uint8) {
	fmt.Printf("[sc]   %x\n", value)
}
