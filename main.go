package main

import (
	"fmt"
	"monty-hall/internal/doors"
	"monty-hall/internal/stats"
)

const LIMIT = 1000000

func main() {
	fmt.Println("+--------------------------------+")
	fmt.Println("| Monthy-Hall Paradox Simulation |")
	fmt.Println("+--------------------------------+")
	fmt.Println("Trying both strategies over", LIMIT, "iterations...")

	// loop
	s := stats.New()
	for s.Iterations = 0; s.Iterations < LIMIT; s.Iterations++ {
		d := doors.New()

		// update door stats
		s.UpdateDoors(d)

		// apply strategies
		stay := d.Stay()
		change := d.Change()

		// update win stats
		if stay {
			s.Wins.Stay++
		}
		if change {
			s.Wins.Change++
		}

		// fmt.Printf("ITERATION: %d/%d, DOORS: %t|%t|%t, STAY: %t, CHANGE: %t\n",
		// 	s.Iterations, LIMIT, d[0], d[1], d[2], stay, change)
	}

	// stats
	fmt.Println()
	fmt.Println("+------------+")
	fmt.Println("| STATISTICS |")
	fmt.Println("+------------+")
	fmt.Println()
	fmt.Println("Total iterations:", s.Iterations)
	fmt.Println()
	fmt.Println("CAR BY DOOR:")
	fmt.Printf("  - A: %d (%.2f%%)\n",
		s.CarDoors[0], (float64(s.CarDoors[0]) / float64(s.Iterations) * 100.0))
	fmt.Printf("  - B: %d (%.2f%%)\n",
		s.CarDoors[1], (float64(s.CarDoors[1]) / float64(s.Iterations) * 100.0))
	fmt.Printf("  - C: %d (%.2f%%)\n",
		s.CarDoors[2], (float64(s.CarDoors[2]) / float64(s.Iterations) * 100.0))
	fmt.Println()
	fmt.Printf("WINNING STRATEGY IS \"%s\":\n", s.Winner())
	fmt.Printf("  - STAY WINS...: %d/%d (%.2f%%)\n",
		s.Wins.Stay, s.Iterations, (float64(s.Wins.Stay) / float64(s.Iterations) * 100.0))
	fmt.Printf("  - CHANGE WINS.: %d/%d (%.2f%%)\n",
		s.Wins.Change, s.Iterations, (float64(s.Wins.Change) / float64(s.Iterations) * 100.0))
	fmt.Println()

	if s.Winner() == "CHANGE" {
		fmt.Println("This confirms the theory, changing the first choice is the best.")
	} else {
		fmt.Println("This disproves the theory, but is imposible so check the program and the computer.")
		fmt.Println("Random number generation its not working properly?.")
	}

	fmt.Println()
}
