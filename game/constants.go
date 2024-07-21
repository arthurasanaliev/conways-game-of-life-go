package game

const (
	SCREEN_WIDTH  = 1000
	SCREEN_HEIGHT = 1000

	WINDOW_WIDTH  = 1000
	WINDOW_HEIGHT = 1000

	CELL_WIDTH  = 10
	CELL_HEIGHT = 10

	UPDATE_INTERVAL = 0.05
)

var (
	NUMBER_OF_DIRECTIONS = 8
	DIRECTIONS_X         = [8]int{1, 1, 1, 0, 0, -1, -1, -1}
	DIRECTIONS_Y         = [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
)
