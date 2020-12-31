package structural

// TODO ...
// Moves elements into tree structures. Composite allows
// clients to treat individual elements and groups
// of elements uniformly.

type Component interface {
	Add(x, y int)
	Remove(x, y int)
	GoBack() []Component
}

type Maze struct {
	rat    string
	cheese string
	nodes  [][]Component
}

func (d *Wall) Add(x, y int) {
	d.nodes[index] = append(d.nodes[index], child)
}

func (d *Wall) WhichNode() string {
	return d.name
}

func (d *Wall) GoBack() []Component {
	// TODO ...
}

type Path struct {
	name string
}

// TODO ...

func NewWall(x int, y int) *Wall {
	return &Wall{
		x: x,
		y: y,
	}
}

func NewPath(x int, y int) *Path {
	return &Path{
		x: x,
		y: y,
	}
}
