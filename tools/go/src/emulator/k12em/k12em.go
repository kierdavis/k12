package main

// Local imports
import (
	"emulator"
)

// Fibonacci program
var fib = []uint16{
	0x1F00, 0x5F01, 0x8C00, 0x0F00, 0x6800, 0x5000, 0xC7FC,
}

func main() {
	em := emulator.New()
	em.Trace = emulator.SimpleTrace
	em.OutputPorts[0] = emulator.DebugOutputPort{}
	em.LoadProgram(fib, 0)

	em.Run(60)
}
