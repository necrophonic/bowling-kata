package bowling_test

import (
	"bowling-kata/bowling"
	"testing"
)

func TestSimpleRoll(t *testing.T) {
	g := bowling.StartGame()
	g.Roll(2)
	score := g.Score()
	if score != 2 {
		t.Errorf("Expected %d, to be 2", score)
	}
}

func TestTooManyRolls(t *testing.T) {
	g := bowling.StartGame()
	for i := 0; i < 22; i++ {
		g.Roll(1)
	}

	if err := g.Roll(1); err == nil {
		t.Errorf("Expected error, but got %v", err)
	}
}

func TestTwoFrameScore(t *testing.T) {
	g := bowling.StartGame()
	g.Roll(2)
	g.Roll(1)
	g.Roll(7)
	g.Roll(1)
	score := g.Score()
	if score != 11 {
		t.Errorf("Expected %d, to be 11", score)
	}
}

func TestStrike(t *testing.T) {
	g := bowling.StartGame()
	g.Roll(10) // Ball 1
	g.Roll(1)  // Ball 3
	g.Roll(7)  // Ball 4
	score := g.Score()
	if score != 26 {
		t.Errorf("Expected %d, to be 26", score)
	}
}

func TestSpare(t *testing.T) {
	g := bowling.StartGame()
	g.Roll(9) // Ball 1
	g.Roll(1) // Ball 2 - spare
	g.Roll(7) // Ball 3
	g.Roll(2) // Ball 4
	score := g.Score()
	if score != 26 {
		t.Errorf("Expected %d, to be 26", score)
	}
}

func TestFinalSpare(t *testing.T) {
	g := bowling.StartGame()
	for i := 0; i < 19; i++ {
		g.Roll(0)
	}
	g.Roll(10) // Spare
	g.Roll(2)
	score := g.Score()
	if score != 14 {
		t.Errorf("Expected %d, to be 14", score)
	}
}

func TestFinalSpareNotEnoughBalls(t *testing.T) {
	g := bowling.StartGame()
	for i := 0; i < 19; i++ {
		g.Roll(0)
	}
	g.Roll(10) // Spare
	// g.Roll(2) // No final ball
	score := g.Score()
	if score != 10 {
		t.Errorf("Expected %d, to be 10", score)
	}
}

func TestPerfectGame(t *testing.T) {
	g := bowling.StartGame()
	for i := 0; i < 12; i++ {
		g.Roll(10)
	}
	score := g.Score()
	if score != 300 {
		t.Errorf("Expected %d, to be 300", score)
	}
}

func TestImperfectGame(t *testing.T) {
	g := bowling.StartGame()
	for i := 0; i < 10; i++ {
		g.Roll(10)
	}
	g.Roll(1)
	g.Roll(1)
	score := g.Score()
	if score != 194 {
		t.Errorf("Expected %d, to be 194", score)
	}
}
