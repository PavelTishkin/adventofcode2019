package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
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
	outputFile := os.Args[2]
	cellsPainted := getNumberOfPaintedCells(input[0])

	fmt.Printf("Part 1: %d\n", cellsPainted)
	paintFromWhite(input[0], outputFile)
	fmt.Printf("Part 2: see %s\n", outputFile)
}

func paintFromWhite(input string, output string) {
	var p = intcode.Program{}
	p.InitMemory(input)

	bot := paintBot{}
	bot.initBot(p)
	bot.paintCell(1)
	bot.run()

	paintedCells := bot.paintedCells
	xOffset, yOffset := getMin(paintedCells)
	adjustArray(&paintedCells, xOffset, yOffset)
	maxX, maxY := getMax(paintedCells)

	convertArrayToPng(maxX, maxY, paintedCells, output)
}

func getMin(cells map[point]int) (int, int) {
	minX := 0
	minY := 0
	for p := range cells {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
	}
	return minX, minY
}

func getMax(cells map[point]int) (int, int) {
	maxX := 0
	maxY := 0
	for p := range cells {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	return maxX + 1, maxY + 1
}

func adjustArray(cells *map[point]int, xOffset int, yOffset int) {
	for p := range *cells {
		p.x -= xOffset
		p.y -= yOffset
	}
}

func convertArrayToPng(width int, height int, cells map[point]int, imagePath string) {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	black := color.RGBA{0, 0, 0, 0xff}
	//tranparent := color.RGBA{0, 0, 0, 0}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, black)
			for p, col := range cells {
				currPoint := point{x: x, y: y}
				if p == currPoint && col == 1 {
					img.Set(x, y, white)
				}
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create(imagePath)
	png.Encode(f, img)
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
