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
	Send
	Receive
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

// Execute the statement, pointed to by idxCurrentStatement.
// Since a statement can influence the RunEnv, we require the RunEnv to be sent as well
func (s *Statement) exec(p *Program) {
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
			p.valueLastPlayed = s.valueLeft.getValue()
		}
	case Recovery:
		if s.valueLeft.getValue() != 0 {
			fmt.Println("exec: isRecovered is true")
			p.state = isRecovered
		} else {
			fmt.Println("exec: isRecovered is still false: leftValue is 0")
		}
	case JumpGreaterZero:
		if s.valueLeft.getValue() != 0 {
			p.idxCurrentStatement += -1 + s.valueRight.getValue()
			fmt.Println("exec: jump to statement: ", p.idxCurrentStatement)
		} else {
			fmt.Println("exec: skip jumping, leftValue is 0")
		}
	case Send:
		fmt.Println("exec: sending value to other program:", s.valueLeft.getValue())
		p.occerencesValueSend++
		p.chanToOtherProgram <- s.valueLeft.getValue()
	case Receive:
		p.state = isReceiving
		fmt.Println("exec: waiting to receive value from otherprogram:", p.chanToUniverse)
    p.chanToUniverse <- struct {
      programId string,
      state stateProgram 
    }{
      programId: p.identifier,
      state: isReceiving,
    }
		// s.valueLeft.setValue(<-p.chanToOtherProgram)
		fmt.Println("exec: value received:", s.valueLeft.getValue())
		p.state = isRunning
	}
}
