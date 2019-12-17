package main

import (
	"fmt"
	"math"
	"os"

	"../intcode"
	"../utils"
)

type point struct {
	x int
	y int
}
type paintBot struct {
	code         intcode.Program
	location     point
	direction    int
	paintedCells map[point]int
}

func main() {
	input := utils.ReadLines(os.Args[1])
	cellsPainted := getNumberOfPaintedCells(input[0])

	fmt.Printf("Part 1: %d\n", cellsPainted)
}

func getNumberOfPaintedCells(input string) int {
	var p = intcode.Program{}
	p.InitMemory(input)

	bot := paintBot{}
	bot.initBot(p)
	bot.run()

	return len(bot.paintedCells)
}

func (bot *paintBot) initBot(p intcode.Program) {
	bot.code = p
	bot.location = point{x: 0, y: 0}
	bot.direction = 0
	bot.paintedCells = make(map[point]int)
	bot.code.Run()
}

func (bot *paintBot) getCellColor() int {
	if cellColor, ok := bot.paintedCells[bot.location]; ok {
		return cellColor
	}
	return 0
}

func (bot *paintBot) paintCell(color int) {
	bot.paintedCells[bot.location] = color
}

func (bot *paintBot) turnLeft() {
	bot.direction = int(math.Mod(float64(bot.direction+270), 360))
}

func (bot *paintBot) turnRight() {
	bot.direction = int(math.Mod(float64(bot.direction+90), 360))
}

func (bot *paintBot) move() {
	switch bot.direction {
	case 0:
		bot.location.y--
	case 90:
		bot.location.x++
	case 180:
		bot.location.y++
	case 270:
		bot.location.x--
	}
}

func (bot *paintBot) run() {
	botCode := &bot.code
	for botCode.IsRunning() {
		currCellColor := int64(bot.getCellColor())
		botCode.PushInput(currCellColor)
		botCode.Continue()
		paintColor := botCode.PopOutput()
		bot.paintCell(int(paintColor))
		turnDirection := botCode.PopOutput()
		if turnDirection == 0 {
			bot.turnLeft()
		} else {
			bot.turnRight()
		}
		bot.move()
	}
}
