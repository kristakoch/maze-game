package structural

// Mazer provides base functionality
// for the  maze to operate.
type Mazer interface {
	AddCheese() string
}

// Cheese ...
type Cheese struct {
}

// NewMaze is the Adapter constructor.
func NewMaze(cheeseType Cheese) Mazer {
	return Maze{cheeseType}
}

// AddGruyere will add gruyere to the maze.
func (c Cheese) AddGruyere() string {
	return "Gruyere!"
}

// Maze embeds the cheese type so that
// we can apply cheese directly from the
// Maze type.
type Maze struct {
	Cheese
}

// AddCheese is an adaptive method.
func (a Maze) AddCheese() string {
	return a.AddGruyere()
}
