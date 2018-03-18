package instruction

import (
	"fmt"
	"testing"
)

// Test string method for instructions which
// have a defined display.
func TestAvailableDisplays(t *testing.T) {
	if PSH.String() != "PSH" {
		ErrorExpectedString("PSH", string(PSH), t)
	}
	if ADD.String() != "ADD" {
		ErrorExpectedString("ADD", string(ADD), t)
	}
	if POP.String() != "POP" {
		ErrorExpectedString("POP", string(POP), t)
	}
	if MOV.String() != "MOV" {
		ErrorExpectedString("MOV", string(MOV), t)
	}
	if HLT.String() != "HLT" {
		ErrorExpectedString("HLT", string(HLT), t)
	}
}

// Test a string method for an instruction
// which is undefined.
//func TestUnavailableDisplays(t *testing.T) {
//	recieved := string(Instruction(10))
//	expected := "Instruction(10)"
//	if expected != recieved {
//		ErrorExpectedString(expected, recieved, t)
//	}
//}

func ErrorExpectedString(expected string, recieved string, t *testing.T) {
	t.Error(fmt.Sprintf("Error: Expected formatted instruction \"%s\", recieved \"%s\".", expected, recieved))
}
