package puzzle

type stateProgram uint

const (
	isRunning stateProgram = iota
	isReceiving
	isRecovered
)

type Program struct {
	identifier          string
	statements          []*Statement                // All the statments, read from the input-file into a slice, so we can jump forwards and backwards
	values              map[string]Value            // All the values, used in statements, index by their name; can be either register or raw number
	idxCurrentStatement int                         // Index of the statement currently being executed
	valueLastPlayed     int                         // Value played by the last, successfully called snd-operation
	state               stateProgram                // Current state of the program, to detect end-state and deadlocks
	occerencesValueSend int                         // Times this program send a value to the other program
	chanToOtherProgram  chan int                    // Channel to send values to the other program
	chanToUniverse      chan struct{ stateProgram } // Channel to send values to the universe, so it can detect endstates
}

// Set default values and sizes for the variables and values
func (p *Program) initProgram(nrStatements int) *Program {
	p.statements = make([]*Statement, nrStatements)
	p.values = make(map[string]Value, nrStatements*2)
	p.state = isRunning
	p.idxCurrentStatement = 0

	return p
}

func (p *Program) execStatements() *Program {
	for p.state == isRunning {
		// fmt.Println("\nexecStatements:")
		// fmt.Println("idxCurrentStatement:", u.idxCurrentStatement)
		// fmt.Println("line:", u.statements[u.idxCurrentStatement].origLine)

		p.statements[p.idxCurrentStatement].exec(p)
		// fmt.Println("Value of register 'a':", u.values["a"].getValue())

		p.idxCurrentStatement += 1
	}

	return p
}
