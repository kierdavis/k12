package emulator

// Standard library imports
import (
    "fmt"
)

// Represents an input port.
type InputPort interface {
    Read() uint8
}

// Represents an output port.
type OutputPort interface {
    Write(uint8)
}

// Dumps all values written to it to standard output.
type DebugOutputPort struct {}

func (DebugOutputPort) Write(x uint8) {
	fmt.Printf("OUT: dec %3d % 3d  hex %02X  char %q\n", x, int8(x), x, x)
}
