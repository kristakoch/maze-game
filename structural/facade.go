package structural

// Facade defines a more precise interface that makes a subsystem easier to use.

package structural

import "strings"

func NewMaze() *Maze {
	return &Maze{
		rat:    &Rat{},
		cheese: &Cheese{},
		board:  &Board{},
	}
}

type Maze struct {
	rat    *Rat
	cheese *Cheese
	board  *Board
}

func (m *Maze) Construct() string {
	result := []string{
		m.cheese.DropCheese(),
		m.rat.DropRat(),
		m.board.ConstructBoard(),
	}
	return strings.Join(result, "\n")
}

type Cheese struct {
}

func (c *Cheese) DropCheese() string {
	return "Gruyere"
}

type DropRat struct {
}

func (r *Rat) DropRat() string {
	return "squeak"
}

type Board struct {
}

func (b *Board) ConstructBoard() string {
	return "board is constructeed"
}