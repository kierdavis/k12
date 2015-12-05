package emulator

// Standard library imports
import (
	"fmt"
)

// A function suitable for assigning to the Emulator.Trace field that simply
// dumps the registers contents and current instruction to standard output.
func SimpleTrace(em *Emulator, inst uint16) {
	var str string

	if inst&0x0800 != 0 {
		switch (inst >> 12) & 0x3 {
		case 0x0, 0x1:
			var b string
			if inst&0x1000 != 0 {
				b = fmt.Sprintf("%02X", inst&0xFF)
			} else {
				b = "b"
			}
			switch (inst >> 8) & 0x7 {
			case 0x0:
				str = "a"
			case 0x1:
				str = "a & " + b
			case 0x2:
				str = "a | " + b
			case 0x3:
				str = "a ^ " + b
			case 0x4:
				str = "a + " + b
			case 0x5:
				str = "a - " + b
			case 0x6:
				str = "a >> 1"
			case 0x7:
				str = b
			}
		case 0x2:
			str = "c"
		case 0x3:
			str = "d"
		}
		switch (inst >> 14) & 0x3 {
		case 0x0:
			str = "mov a <- " + str
		case 0x1:
			str = "mov b <- " + str
		case 0x2:
			str = "mov c <- " + str
		case 0x3:
			str = "mov d <- " + str
		}
	} else {
		switch (inst >> 12) & 0xF {
		case 0x0:
			str = "inc"
		case 0x1:
			str = fmt.Sprintf("in %d", inst&0x7)
		case 0x2:
			str = "ldi"
		case 0x3:
			str = fmt.Sprintf("ldd %04X", inst&0xFF)
		case 0x4:
			str = "dec"
		case 0x5:
			str = fmt.Sprintf("out %d", inst&0x7)
		case 0x6:
			str = "std"
		case 0x7:
			str = fmt.Sprintf("sti %04X", inst&0xFF)
		case 0x8, 0x9, 0xA, 0xB:
			if inst&0x1000 != 0 {
				str = fmt.Sprintf("skip (a - %02X) if ", inst&0xFF)
			} else {
				str = "skip (a - b) if "
			}
			if inst&0x2000 != 0 {
				str += simpleTraceFlagsNeg[(inst>>8)&0x7]
			} else {
				str += simpleTraceFlagsPos[(inst>>8)&0x7]
			}
		case 0xC:
			rel := inst & 0x07FF
			if rel&0x0400 != 0 {
				rel -= 0x0800
			}
			str = fmt.Sprintf("rjmp %+05X (to %04X)", int16(rel), em.PC+rel)
		case 0xD:
			rel := inst & 0x07FF
			if rel&0x0400 != 0 {
				rel -= 0x0800
			}
			str = fmt.Sprintf("rcall %+05X (to %04X)", int16(rel), em.PC+rel)
		case 0xE:
			str = fmt.Sprintf("ljmp (to %02X%02X)", em.C, em.D)
		case 0xF:
			str = "halt"
		}
	}

	fmt.Printf("[%02X %02X %02X %02X %04X] %04X %s\n", em.A, em.B, em.C, em.D, em.PC, inst, str)
}

var simpleTraceFlagsPos = [8]string{
	"zero/equal",
	"negative",
	"borrow",
	"overflow",
	"unsigned lt",
	"unsigned le",
	"signed lt",
	"signed le",
}

var simpleTraceFlagsNeg = [8]string{
	"not zero/equal",
	"not negative",
	"not borrow",
	"not overflow",
	"unsigned ge",
	"unsigned gt",
	"signed ge",
	"signed gt",
}
