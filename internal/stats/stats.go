package stats

import "monty-hall/internal/doors"

type Stats struct {
	CarDoors   [3]int
	Iterations int
	Wins       struct {
		Stay   int
		Change int
	}
	Looses struct {
		Stay   int
		Change int
	}
}

func New() *Stats {
	return new(Stats)
}

func (s *Stats) UpdateDoors(d *doors.Doors) {
	if d[0] {
		s.CarDoors[0]++
		return
	}
	if d[1] {
		s.CarDoors[1]++
		return
	}
	s.CarDoors[2]++
}

func (s *Stats) Winner() string {
	if s.Wins.Stay > s.Wins.Change {
		return "STAY"
	}
	if s.Wins.Change > s.Wins.Stay {
		return "CHANGE"
	}
	return "TIED"
}
