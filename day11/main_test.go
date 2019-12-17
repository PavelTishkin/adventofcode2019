package main

import (
	"testing"

	"../intcode"
)

func TestGetCellColor(t *testing.T) {
	bot := paintBot{}
	p := intcode.Program{}
	p.InitMemory("99")
	bot.initBot(p)

	bot.paintCell(1)
	cellColor := bot.getCellColor()
	if cellColor != 1 {
		t.Errorf("cellColor = %d, expected 1", cellColor)
	}

	bot.paintCell(0)
	cellColor = bot.getCellColor()
	if cellColor != 0 {
		t.Errorf("cellColor = %d, expected 1", cellColor)
	}

	bot.move()
	cellColor = bot.getCellColor()
	if cellColor != 0 {
		t.Errorf("cellColor = %d, expected 1", cellColor)
	}
}

func TestPaintCell(t *testing.T) {
	bot := paintBot{}
	p := intcode.Program{}
	p.InitMemory("99")
	bot.initBot(p)

	bot.paintCell(1)
	if bot.paintedCells[bot.location] != 1 {
		t.Errorf("cellColor = %d, expected 1", bot.paintedCells[bot.location])
	}
}

func TestTurnLeft(t *testing.T) {
	bot := paintBot{}
	p := intcode.Program{}
	p.InitMemory("99")
	bot.initBot(p)

	bot.turnLeft()
	expected := 270
	if bot.direction != expected {
		t.Errorf("bot.direction = %d, expected %d", bot.direction, expected)
	}

	bot.turnLeft()
	bot.turnLeft()
	bot.turnLeft()
	expected = 0
	if bot.direction != expected {
		t.Errorf("bot.direction = %d, expected %d", bot.direction, expected)
	}
}

func TestTurnRight(t *testing.T) {
	bot := paintBot{}
	p := intcode.Program{}
	p.InitMemory("99")
	bot.initBot(p)

	bot.turnRight()
	expected := 90
	if bot.direction != expected {
		t.Errorf("bot.direction = %d, expected %d", bot.direction, expected)
	}

	bot.turnRight()
	bot.turnRight()
	bot.turnRight()
	expected = 0
	if bot.direction != expected {
		t.Errorf("bot.direction = %d, expected %d", bot.direction, expected)
	}
}

func TestMove(t *testing.T) {
	bot := paintBot{}
	p := intcode.Program{}
	p.InitMemory("99")
	bot.initBot(p)

	bot.move()
	expected := point{x: 0, y: -1}
	if bot.location != expected {
		t.Errorf("bot.location = %v, expected %v", bot.location, expected)
	}

	bot.turnLeft()
	bot.move()
	expected = point{x: -1, y: -1}
	if bot.location != expected {
		t.Errorf("bot.location = %v, expected %v", bot.location, expected)
	}

	bot.turnLeft()
	bot.move()
	expected = point{x: -1, y: 0}
	if bot.location != expected {
		t.Errorf("bot.location = %v, expected %v", bot.location, expected)
	}

	bot.turnLeft()
	bot.move()
	expected = point{x: 0, y: 0}
	if bot.location != expected {
		t.Errorf("bot.location = %v, expected %v", bot.location, expected)
	}
}
