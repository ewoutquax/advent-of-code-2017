package puzzle

import (
	"strconv"
	"strings"
)

type Universe struct {
	statements          []*Statement     // All the statments, read from the input-file in a slice, so we can jump forwards and backwards
	values              map[string]Value // All the values, used in statements, index by their name; can be either register or raw number
	idxCurrentStatement int              // Index of the statement currently being executed
	valueLastPlayed     int              // value played by the last, successfully called snd-operation
	isRecovered         bool             // Has the recover-instruction been executed
}

// Set default values and sizes for the variables and values
func (u *Universe) initUniverse(nrStatements int) *Universe {
	u.statements = make([]*Statement, nrStatements)
	u.values = make(map[string]Value, nrStatements*2)
	u.isRecovered = false
	u.idxCurrentStatement = 0

	return u
}

func (u *Universe) execStatements() *Universe {
	for !u.isRecovered {
		fmt.Println("\nexecStatements:")
		fmt.Println("idxCurrentStatement:", u.idxCurrentStatement)
		fmt.Println("line:", u.statements[u.idxCurrentStatement].origLine)

		u.statements[u.idxCurrentStatement].exec(u)
		fmt.Println("Value of register 'a':", u.values["a"].getValue())

		u.idxCurrentStatement += 1
	}

	return u
}

func FindLastPlayedValue(lines []string) int {
	universe := parseInput(lines)

	return universe.execStatements().valueLastPlayed
}

func parseInput(lines []string) (universe Universe) {
	var parts []string
	universe.initUniverse(len(lines))

	for idx, line := range lines {
		parts = strings.Split(line, " ")

		s := Statement{
			origLine:  line,
			Operation: convStringToOperation(parts[0]),
		}

		if value, exists := universe.values[parts[1]]; exists {
			s.valueLeft = value
		} else {
			universe.values[parts[1]] = convStringToValue(parts[1])
			s.valueLeft = universe.values[parts[1]]
		}

		if len(parts) == 3 {
			if value, exists := universe.values[parts[2]]; exists {
				s.valueRight = value
			} else {
				universe.values[parts[2]] = convStringToValue(parts[2])
				s.valueRight = universe.values[parts[2]]
			}
		}

		universe.statements[idx] = &s
	}

	return
}

func convStringToOperation(part string) Operation {
	return map[string]Operation{
		"set": Set,
		"add": Add,
		"mul": Multiply,
		"mod": Modulo,
		"snd": Sound,
		"rcv": Recovery,
		"jgz": JumpGreaterZero,
	}[part]
}

func convStringToValue(part string) Value {
	number, err := strconv.Atoi(part)

	if err == nil {
		return &Number{value: number}
	} else {
		return &Register{value: number}
	}
}
