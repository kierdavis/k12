package emulator

// Standard library imports
import (
	"encoding/binary"
)

// Represents the state of a K12 processor.
type Emulator struct {
	// Data registers
	A, B, C, D uint8

	// Program counter
	PC uint16

	// Flag indicating whether processor is running or halted
	Halted bool

	// Number of elapsed clock cycles since last reset
	Cycles uint

	// Memories
	DataMem *[65536]uint8
	ProgMem *[65536]uint16

	// I/O ports
	InputPorts  [8]InputPort
	OutputPorts [8]OutputPort

	// Trace function; assign a non-nil function to this, and it will be executed
	// before each instruction is executed.
	Trace func(em *Emulator, inst uint16)
}

func New() (em *Emulator) {
	em = new(Emulator)
	em.DataMem = new([65536]uint8)
	em.ProgMem = new([65536]uint16)
	return em
}

// Copy 'data' into program memory starting at address 'addr'.
func (em *Emulator) LoadProgram(data []uint16, addr uint16) {
	copy(em.ProgMem[addr:], data)
}

// Copy 'data' into program memory starting at address 'addr'. 'byteOrder'
// determines whether the 16-bit words in 'data' are stored in big-endian or
// little-endian form.
func (em *Emulator) LoadProgramBytes(data []uint8, addr uint16, byteOrder binary.ByteOrder) {
	for i := 0; i+1 < len(data); i += 2 {
		em.ProgMem[addr] = byteOrder.Uint16(data[i:])
		addr++
	}
}

// Sets all registers and the cycle counter to zero, and sets the processor to the
// running (not halted) state. Does not affect either of the memories.
func (em *Emulator) ResetRegisters() {
	em.A = 0
	em.B = 0
	em.C = 0
	em.D = 0
	em.PC = 0
	em.Halted = false
	em.Cycles = 0
}

// Reset the contents of data memory.
func (em *Emulator) ResetDataMemory() {
	for i := range em.DataMem {
		em.DataMem[i] = 0
	}
}

// Run for the given number of cycles, or until the emulator becomes halted. If
// it is already halted, return immediately.
func (em *Emulator) Run(cycles uint) {
	for !em.Halted && cycles != 0 {
		em.Step()
		cycles--
	}
}

// Run the emulator until the emulator becomes halted. If it is already halted,
// return immediately.
func (em *Emulator) RunUntilHalt() {
	for !em.Halted {
		em.Step()
	}
}

// Execute a single instruction.
func (em *Emulator) Step() {
	inst := em.ProgMem[em.PC]

	if em.Trace != nil {
		em.Trace(em, inst)
	}

	var handler func(*Emulator, uint16)

	if inst&0x0800 != 0 {
		handler = (*Emulator).doMov
	} else {
		switch (inst >> 12) & 0xF {
		case 0x0:
			handler = (*Emulator).doInc
		case 0x1:
			handler = (*Emulator).doIn
		case 0x2:
			handler = (*Emulator).doLdi
		case 0x3:
			handler = (*Emulator).doLdd
		case 0x4:
			handler = (*Emulator).doDec
		case 0x5:
			handler = (*Emulator).doOut
		case 0x6:
			handler = (*Emulator).doSti
		case 0x7:
			handler = (*Emulator).doStd
		case 0x8, 0x9, 0xA, 0xB:
			handler = (*Emulator).doSkip
		case 0xC:
			handler = (*Emulator).doRjmp
		case 0xD:
			handler = (*Emulator).doRcall
		case 0xE:
			handler = (*Emulator).doLjmp
		case 0xF:
			handler = (*Emulator).doHalt
		}
	}

	handler(em, inst)
}

// Data move.
func (em *Emulator) doMov(inst uint16) {
	var data uint8

	// Source
	switch (inst >> 12) & 0x3 {
	case 0x0, 0x1:
		// ALU
		a := em.A
		var b uint8
		if inst&0x1000 != 0 {
			b = uint8(inst)
		} else {
			b = em.B
		}

		switch (inst >> 8) & 0x7 {
		case 0x0:
			data = a
		case 0x1:
			data = a & b
		case 0x2:
			data = a | b
		case 0x3:
			data = a ^ b
		case 0x4:
			data = a + b
		case 0x5:
			data = a - b
		case 0x6:
			data = (a >> 1) | (a & 0x80)
		case 0x7:
			data = b
		}

	case 0x2:
		data = em.C
	case 0x3:
		data = em.D
	}

	// Destination
	switch (inst >> 14) & 0x3 {
	case 0x0:
		em.A = data
	case 0x1:
		em.B = data
	case 0x2:
		em.C = data
	case 0x3:
		em.D = data
	}

	em.PC++
}

// Increment CD.
func (em *Emulator) doInc(inst uint16) {
	em.D++
	if em.D == 0x00 { // overflow
		em.C++
	}
	em.PC++
}

// Read from IO port.
func (em *Emulator) doIn(inst uint16) {
	port := em.InputPorts[inst&0x7]
	em.A = port.Read()
	em.PC++
}

// Load indirect.
func (em *Emulator) doLdi(inst uint16) {
	addr := (uint16(em.C) << 8) | uint16(em.D)
	em.A = em.DataMem[addr]
	em.PC++
}

// Load direct.
func (em *Emulator) doLdd(inst uint16) {
	addr := inst & 0x00FF
	em.A = em.DataMem[addr]
	em.PC++
}

// Decrement CD.
func (em *Emulator) doDec(inst uint16) {
	em.D--
	if em.D == 0xFF { // underflow
		em.C--
	}
	em.PC++
}

// Write to IO port.
func (em *Emulator) doOut(inst uint16) {
	port := em.OutputPorts[inst&0x7]
	port.Write(em.A)
	em.PC++
}

// Store indirect.
func (em *Emulator) doSti(inst uint16) {
	addr := (uint16(em.C) << 8) | uint16(em.D)
	em.DataMem[addr] = em.A
	em.PC++
}

// Store direct.
func (em *Emulator) doStd(inst uint16) {
	addr := inst & 0x00FF
	em.DataMem[addr] = em.A
	em.PC++
}

// Conditional skip.
func (em *Emulator) doSkip(inst uint16) {
	// Get operands
	a := em.A
	var b uint8
	if inst&0x1000 != 0 {
		b = uint8(inst)
	} else {
		b = em.B
	}

	// Perform subtraction using a + ^b + 1
	b = ^b
	result := uint16(a) + uint16(b) + 1

	// Compute flags
	zero := (result & 0xFF) == 0
	negative := (result & 0x80) != 0
	borrow := (result & 0x100) == 0
	overflow := ((a ^ uint8(result)) & (b ^ uint8(result)) & 0x80) != 0
	slt := negative != overflow

	var flag bool

	switch (inst >> 8) & 0x7 {
	case 0x0:
		flag = zero
	case 0x1:
		flag = negative
	case 0x2, 0x4:
		flag = borrow
	case 0x3:
		flag = overflow
	case 0x5:
		flag = borrow || zero
	case 0x6:
		flag = slt
	case 0x7:
		flag = slt || zero
	}

	if inst&0x2000 != 0 {
		flag = !flag
	}

	if flag {
		em.PC += 2
	} else {
		em.PC++
	}
}

// Relative jump.
func (em *Emulator) doRjmp(inst uint16) {
	rel := inst & 0x07FF

	// sign-extend from 11 to 16 bits
	if rel&0x0400 != 0 {
		rel -= 0x0800
	}

	em.PC += rel
}

// Relative call.
func (em *Emulator) doRcall(inst uint16) {
	em.C = uint8(em.PC >> 8)
	em.D = uint8(em.PC)
	em.doRjmp(inst)
}

// Long (indirect) jump.
func (em *Emulator) doLjmp(inst uint16) {
	em.PC = (uint16(em.C) << 8) | uint16(em.D)
}

// Halt.
func (em *Emulator) doHalt(inst uint16) {
	em.Halted = true
	em.PC++
}
