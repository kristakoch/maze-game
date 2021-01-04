package pattern

/*

Pattern:
memento

Premise:
Maze visitors are offered a few chances to push a button at a computer to
generate and return to one checkpoint. While they don't know exactly what the
game state is made up of (something only the computer knows), they can save
and restore their state.

Participants:
- Concrete memento: checkpoint, which stores a private state snapshot.
  Exporting memento lets outside parties use them without having to make
  the originator's state public, which would break encapsulation, causing changes
  to the originator ripple out and require changes in things using it.
- Concrete originator: thing with private state that can generate mementos and
  access private fields and methods on the memento.
- Caretaker: delegates when state is saved and restored. Here it's the maze
  visitor, who can create and restore checkpoints but knows nothing of the
  actual fields involved (like whether an audience likes them).

*/

type gameState struct {
	hungerLevel     int
	keysFound       int
	furthestPercent float64
	audienceRating  int
}

type checkpointMemento struct {
	snapshot gameState
}

type mazeSystemOriginator struct {
	state gameState
	m     *checkpointMemento
}

func (o *mazeSystemOriginator) restore(m *checkpointMemento) {
	o.state = m.snapshot
}

func (o *mazeSystemOriginator) save() *checkpointMemento {
	if o.m == nil {
		o.m = new(checkpointMemento)
	}
	o.m.createCheckpointMemento(o.state)
	return o.m
}

func (m *checkpointMemento) createCheckpointMemento(state gameState) {
	m.snapshot = state
}

type mazeVisitorCaretaker struct {
	o *mazeSystemOriginator
	m *checkpointMemento
}

func (c *mazeVisitorCaretaker) returnToCheckpoint() {
	c.o.restore(c.m)
}

func (c *mazeVisitorCaretaker) createCheckpoint() {
	c.m = c.o.save()
}
