package puzzle

import "fmt"

type Operation uint

const (
	Set Operation = iota
	Add
	Multiply
	Modulo
	Sound
	Recovery
	JumpGreaterZero
)

// Interface to enable Statements to reference Registers and Number equally
type Value interface {
	setValue(int)
	getValue() int
}

type Number struct {
	value int // Raw value hold by this struct. Placed in a struct for easy reference from the statements
}

func (n Number) getValue() int {
	return n.value
}

func (n *Number) setValue(_ int) {
	panic("Can not set the value of a raw number")
}

type Register struct {
	value int // Numeric value hold by the registry
}

func (r Register) getValue() int {
	return r.value
}

func (r *Register) setValue(value int) {
	r.value = value
}

type Statement struct {
	Operation        // The operation this statement needs to execute
	valueLeft  Value // Left value used by the operation
	valueRight Value // (optional) right value used by the operation
	origLine   string
}

func (s *Statement) exec(u *Universe) {
	switch s.Operation {
	case Set:
		s.valueLeft.setValue(s.valueRight.getValue())
	case Add:
		s.valueLeft.setValue(s.valueLeft.getValue() + s.valueRight.getValue())
	case Multiply:
		s.valueLeft.setValue(s.valueLeft.getValue() * s.valueRight.getValue())
	case Modulo:
		s.valueLeft.setValue(s.valueLeft.getValue() % s.valueRight.getValue())
	case Sound:
		if s.valueLeft.getValue() != 0 {
			fmt.Println("exec: playing sound:", s.valueLeft.getValue())
			u.valueLastPlayed = s.valueLeft.getValue()
		}
	case Recovery:
		if s.valueLeft.getValue() != 0 {
			fmt.Println("exec: isRecovered is true")
			u.isRecovered = true
		} else {
			fmt.Println("exec: isRecovered is still false: leftValue is 0")
		}
	case JumpGreaterZero:
		if s.valueLeft.getValue() != 0 {
			u.idxCurrentStatement += -1 + s.valueRight.getValue()
			fmt.Println("exec: jump to statement: ", u.idxCurrentStatement)
		} else {
			fmt.Println("exec: skip jumping, leftValue is 0")
		}
	}
}
