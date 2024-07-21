package game

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	grid       [][]bool
	lastUpdate time.Time
}

func NewGame() *Game {
	grid := make([][]bool, SCREEN_HEIGHT/CELL_HEIGHT)
	for i := range grid {
		grid[i] = make([]bool, SCREEN_WIDTH/CELL_WIDTH)
	}
	// TODO: make a better cell input
	grid[len(grid)/2-1][len(grid[0])/2-1] = true
	grid[len(grid)/2-1][len(grid[0])/2] = true
	grid[len(grid)/2-1][len(grid[0])/2+1] = true
	grid[len(grid)/2][len(grid[0])/2-1] = true
	grid[len(grid)/2][len(grid[0])/2+1] = true
	grid[len(grid)/2+1][len(grid[0])/2-1] = true
	grid[len(grid)/2+1][len(grid[0])/2+1] = true
	return &Game{
		grid: grid,
	}
}

func (g *Game) Update() error {
	currentTime := time.Now()
	if currentTime.Sub(g.lastUpdate).Seconds() >= UPDATE_INTERVAL {
		g.lastUpdate = currentTime

		out := make([][]bool, len(g.grid))
		for i := range out {
			out[i] = make([]bool, len(g.grid[i]))
		}
		for y, rows := range g.grid {
			for x := range rows {
				count := countNeighbors(g.grid, x, y)
				if g.grid[y][x] {
					if count == 2 || count == 3 {
						out[y][x] = true
					}
				} else {
					if count == 3 {
						out[y][x] = true
					}
				}
			}
		}
		g.grid = out
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	yellowColor := color.RGBA{255, 255, 0, 255}
	for y, rows := range g.grid {
		for x, _ := range rows {
			if g.grid[y][x] {
				vector.DrawFilledRect(screen, float32(x*CELL_WIDTH), float32(y*CELL_HEIGHT), float32(CELL_WIDTH), float32(CELL_HEIGHT), yellowColor, false)
			}
		}
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func countNeighbors(grid [][]bool, x, y int) int {
	count := 0
	for i := 0; i < NUMBER_OF_DIRECTIONS; i++ {
		row := y + DIRECTIONS_Y[i]
		col := x + DIRECTIONS_X[i]
		if row < 0 || row >= len(grid) {
			continue
		}
		if col < 0 || col >= len(grid[0]) {
			continue
		}
		if grid[row][col] {
			count += 1
		}
	}
	return count
}
