package pattern

/*

Pattern:
state

Premise:
The maze can have different modes: easy, medium, hard, and potentially more in
the future. Each mode shares a set of basic actions which can differ based on
the mode the maze is currently in.

Participants:
- State: interface for the basic operations states perform.
- Concrete states: implementations of different states, which will in this
  case implement the operations at the different difficulty levels
- Client: triggers the basic operations, in this case the visitor interacting
  with the maze as they move through it

*/

type mazeState interface {
	shiftWalls()
	releaseEnemy()
	darken()
}

type maze struct {
	state mazeState
}

func (m *maze) changeState(state mazeState) {
	m.state = state
}

func (m *maze) shiftWalls() {
	m.state.shiftWalls()
}
func (m *maze) releaseEnemy() {
	m.state.releaseEnemy()
}
func (m *maze) darken() {
	m.state.darken()
}

type easyMazeState struct{}

func (s *easyMazeState) shiftWalls() {
	// shift a few walls around
}
func (s *easyMazeState) releaseEnemy() {
	// release a fast roomba with a weak taser attached
}
func (s *easyMazeState) darken() {
	// reduce light by 20 percent
}

type mediumMazeState struct{}

func (s *mediumMazeState) shiftWalls() {
	// remove some walls and add some new ones
}
func (s *mediumMazeState) releaseEnemy() {
	// release a volunteer in a monster costume
}
func (s *mediumMazeState) darken() {
	// reduce light by 70 percent
}

type difficultMazeState struct{}

func (s *difficultMazeState) shiftWalls() {
	// redo the maze walls entirely
}
func (s *difficultMazeState) releaseEnemy() {
	// release a minotaur
}
func (s *difficultMazeState) darken() {
	// reduce light by 100 percent
}
