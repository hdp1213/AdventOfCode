package day06

import (
	"strings"
	"testing"
)

// TestTransferBetweenOrbits tests the number of orbital transfers needed to travel from an origin to a destination
func TestTransferBetweenOrbits(t *testing.T) {
	r := strings.NewReader(`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN
`)

	preOrbits, err := readPreOrbits(r)
	if err != nil {
		t.Error(err)
	}

	orbits := processOrbits(preOrbits)

	you, _ := orbits.get(YOU)
	san, _ := orbits.get(SAN)

	transfers := transferBetweenOrbits(you, san)

	if transfers != 4 {
		t.Errorf("expected transfers == %d, got %d\n", 4, transfers)
	}
}
