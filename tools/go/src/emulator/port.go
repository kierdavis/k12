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
type DebugOutputPort struct{}

func (DebugOutputPort) Write(x uint8) {
	fmt.Printf("OUT: dec %3d % 3d  hex %02X  char %q\n", x, int8(x), x, x)
}

// Implements an 8-bit stack.
type Stack struct {
	Data []uint8
}

func (stack *Stack) Read() (x uint8) {
	n := len(stack.Data)
	if n > 0 {
		x = stack.Data[n-1]
		stack.Data = stack.Data[:n-1]
	}
	return x
}

func (stack *Stack) Write(x uint8) {
	stack.Data = append(stack.Data, x)
}
