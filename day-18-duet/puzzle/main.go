package puzzle

import (
	"strconv"
	"strings"
)

type CompilerVersion uint

const (
	v1 CompilerVersion = iota
	v2
)

type Universe struct {
	program1 Program
	program2 Program
}

func FindLastPlayedValue(lines []string) int {
	program := parseInput(lines, v1)

	return program.execStatements().valueLastPlayed
}

func FindOccurencesProgram1Sends(lines []string) int {
	program := parseInput(lines, v2)

	universe := Universe{
		program1: program,
		program2: program,
	}

	universe.program1.chanToOtherProgram = make(chan int)
	universe.program2.chanToOtherProgram = make(chan int)

	go universe.program1.execStatements()
	go universe.program2.execStatements()

	var endState bool = false
	for !endState {
		endState = universe.program1.state == isRecovered ||
			universe.program1.state == isReceiving && universe.program2.state != isRunning ||
			universe.program2.state == isReceiving && universe.program1.state != isRunning
	}

	return universe.program1.occerencesValueSend
}

func parseInput(lines []string, version CompilerVersion) (program Program) {
	var parts []string
	program.initProgram(len(lines))

	for idx, line := range lines {
		parts = strings.Split(line, " ")

		s := Statement{
			origLine:  line,
			Operation: convStringToOperation(parts[0], version),
		}

		if value, exists := program.values[parts[1]]; exists {
			s.valueLeft = value
		} else {
			program.values[parts[1]] = convStringToValue(parts[1])
			s.valueLeft = program.values[parts[1]]
		}

		if len(parts) == 3 {
			if value, exists := program.values[parts[2]]; exists {
				s.valueRight = value
			} else {
				program.values[parts[2]] = convStringToValue(parts[2])
				s.valueRight = program.values[parts[2]]
			}
		}

		program.statements[idx] = &s
	}

	return
}

func convStringToOperation(part string, version CompilerVersion) (op Operation) {
	switch version {
	case v1:
		op = map[string]Operation{
			"set": Set,
			"add": Add,
			"mul": Multiply,
			"mod": Modulo,
			"snd": Sound,
			"rcv": Recovery,
			"jgz": JumpGreaterZero,
		}[part]
	case v2:
		op = map[string]Operation{
			"set": Set,
			"add": Add,
			"mul": Multiply,
			"mod": Modulo,
			"snd": Send,
			"rcv": Receive,
			"jgz": JumpGreaterZero,
		}[part]
	}

	return op
}

func convStringToValue(part string) Value {
	number, err := strconv.Atoi(part)

	if err == nil {
		return &Number{value: number}
	} else {
		return &Register{value: number}
	}
}
