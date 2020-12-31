package structural

// Bridge wll decouple an abstraction from its
// implementation so that the two can vary
// independently.

type Maze interface {
	DropCheese() string
}

type Animal interface {
	Squeak() string
}

type Maze struct {
	animal Animal
}

func NewMaze(animal Animal) *Maze {
	return &Maze{
		animal: animal,
	}
}

func (m *Maze) DropCheese() string {
	return m.animal.Squeak()
}

type Rat struct {
}

func (r *Rat) Squeak() string {
	return "squeak squeak"
}

type GuineaPig struct {
}

func (g *GuineaPig) Squeak() string {
	return "I don't like cheese"
}
