package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLowestDelayForUncaught(t *testing.T) {
	universe := parseInput(input())

	assert.Equal(t, 10, universe.findLowestDelayForUncaught())
}
