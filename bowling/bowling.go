package bowling

import "errors"

// Game represents a bowling game for a single bowler
type Game struct {
	Ball  int   // Current ball being rolled
	Balls []int // All ball scores
}

// StartGame creates a new empty game
func StartGame() *Game {
	return &Game{
		// Declare the whole game. This stops us having to check how many
		// times we've bowled if we want to calculate the score before
		// the actual end of the game
		Balls: make([]int, 22),
	}
}

// Roll represents an individual roll
func (g *Game) Roll(pins int) error {

	// If we've had all our rolls, disregard this one
	// In Go we'd normally return an error but the Kata didn't
	// specify one on it's
	if g.Ball > 21 {
		return errors.New("Too many rolls")
	}

	g.Balls[g.Ball] = pins

	// If a strike then skip the second ball of the frame
	// It's a strike is pins==10 and it's the first ball
	// of the frame
	// Also don't skip if these are the last bonus rolls on
	// the final frame (balls < 20)
	if g.Ball%2 == 0 && pins == 10 && g.Ball < 20 {
		g.Ball++
	}
	g.Ball++

	return nil
}

// Score returns the current game score
func (g *Game) Score() int {
	total := 0
	strike, spare := false, false

	for fn := 1; fn <= 10; fn++ {
		frame := g.Balls[(fn*2)-2 : fn*2]
		score := frame[0] + frame[1]

		if strike {
			score *= 2
			strike = false
		}

		if spare {
			score += frame[0]
			spare = false
		}

		switch {
		case frame[0] == 10:
			strike = true
		case score == 10:
			spare = true
		}

		if fn == 10 && (strike || spare) {
			switch {
			case strike:
				if g.Balls[20]+g.Balls[21] == 20 {
					// Both strikes so perfect game - auto 300
					return 300
				}
				score += (g.Balls[20] + g.Balls[21]) * 2
			case spare:
				score += g.Balls[20] * 2
			}
		}

		total += score
	}
	return total
}
