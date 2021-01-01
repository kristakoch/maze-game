package main

//TypeofFactories provides an interface containing the Bombed and Enchanted factories
type TypeofFactories interface {
	BombedMazeFactory(maze) AbstractBomb
	EnchantedMazeFactory(maze) AbstractEnchantement
}

//AbstractBomb ...
type AbstractBomb interface {
	DestroyWall() []string
}

//AbstractEnchantement ...
type AbstractEnchantement interface {
	EnchantedWall() []string
}

//AbstractFactory provides an interface containing the CreateMaze method
type AbstractFactory interface {
	CreateMaze()
}

func (f *BombedMazeFactory) createMaze() maze {
	return &BombedMazeFactory{}
}

func (f *EnchangedMazeFactory) createMaze() maze {
	return &EnchangedMazeFactory{}
}

type maze struct {
	room1 []string
	room2 []string
	aDoor string
}

//BombedMazeFactory ...
type BombedMazeFactory struct {
}

//EnchangedMazeFactory ...
type EnchangedMazeFactory struct {
}

func main() {

}
