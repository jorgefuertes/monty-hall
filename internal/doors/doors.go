package doors

import (
	"math/rand"
	"time"
)

type Doors [3]bool

func New() *Doors {
	d := new(Doors)
	rand.Seed(time.Now().UnixNano())
	d[rand.Intn(3)] = true
	return d
}

// Stay strategy
func (d *Doors) Stay() bool {
	rand.Seed(time.Now().UnixNano())
	// it stays so he wins if car is in his choosen door
	// and fails if not and should win in a ~33% of times
	return d[rand.Intn(3)]
}

// Change strategy
func (d *Doors) Change() bool {
	// first choice
	choice := rand.Intn(3)

	// presenter shows a goat in one of the other doors
	var pChoice int
	for i, v := range d {
		if i == choice {
			// can't open competitor's door
			continue
		}
		if v {
			// can't open car's door
			continue
		}
		// presenter's door
		pChoice = i
		break
	}

	// second choice
	var sChoice int
	for i, _ := range d {
		if i == choice {
			// can't open first door
			continue
		}
		if i == pChoice {
			// can't open presenter's door
			continue
		}
		// second choice competitor's door
		sChoice = i
		break
	}

	return d[sChoice]
}
