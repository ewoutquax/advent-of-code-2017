package puzzle

type Player struct {
	currentRow     int        // The row on which the player currently is
	hitByScannners []*Scanner // List of all the scanners that have hit the player
}

func (p *Player) tripSeverity() (severity int) {
	for _, scanner := range p.hitByScannners {
		severity += scanner.row * scanner.maxDepth
	}

	return
}

type Universe struct {
	*Player                   // Our player, traveling through this hostile universe
	maxRows  int              // What is the highest row, containing a scanner. This indicates how far the player needs to travel
	scanners map[int]*Scanner // List of all scanners in this universe, indexed by their row number
}

func (u *Universe) initUniverse(nrScanners int) *Universe {
	u.Player = &Player{}
	u.scanners = make(map[int]*Scanner, nrScanners)

	return u
}

func (u Universe) movePlayerToOtherSide() {
	for u.Player.currentRow = 0; u.Player.currentRow <= u.maxRows; u.Player.currentRow++ {
		if scanner, exists := u.scanners[u.Player.currentRow]; exists && scanner.currentLayer == 0 {
			u.Player.hitByScannners = append(u.Player.hitByScannners, u.scanners[u.Player.currentRow])
		}
		for _, scanner := range u.scanners {
			scanner.move()
		}
	}
}

func (u Universe) findLowestDelayForUncaught() int {
	var delay int
	var willCatchPlayer bool = true
	var scanners []*Scanner = make([]*Scanner, len(u.scanners))

	idx := 0
	for _, scanner := range u.scanners {
		scanners[idx] = scanner
		idx++
	}

	for delay = 0; willCatchPlayer; delay++ {
		willCatchPlayer = anyScannerWillCatchPlayer(scanners, delay)
	}

	return delay - 1
}

func anyScannerWillCatchPlayer(scanners []*Scanner, delay int) bool {
	if len(scanners) == 0 {
		return false
	}

	return scanners[0].willCatchPlayerForDelay(delay) || anyScannerWillCatchPlayer(scanners[1:], delay)
}
