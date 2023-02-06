package puzzle

type Direction uint

const (
	ScannerDirectionUp Direction = iota
	ScannerDirectionDown
)

type Scanner struct {
	row          int // On which vertical row is this scanner located
	maxDepth     int // How many layers does this scanner have
	currentLayer int // Which layer is the scanner scanning
	Direction        // Enum, which tells if the scanner is going up or down
}

// Move the scanned layer up or down, depending on the direction of the layer.
// When the laser hits the upper or lower boundary, then reset the direction accordingly.
func (s *Scanner) move() *Scanner {
	if s.currentLayer == 0 {
		s.Direction = ScannerDirectionDown
	}
	if s.currentLayer == s.maxDepth-1 {
		s.Direction = ScannerDirectionUp
	}

	switch s.Direction {
	case ScannerDirectionUp:
		s.currentLayer--
	case ScannerDirectionDown:
		s.currentLayer++
	}

	return s
}

// Return true, when the scanner will hit layer 0 when the player visits this row
// upon starting after the given delay
func (s *Scanner) willCatchPlayerForDelay(delay int) bool {
	return (s.row+delay)%((s.maxDepth-1)*2) == 0
}

func newScanner(row int, depth int) Scanner {
	return Scanner{
		row:          row,
		maxDepth:     depth,
		currentLayer: 0,
		Direction:    ScannerDirectionDown,
	}
}
