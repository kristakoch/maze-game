package pattern

/*

Pattern:
command

Premise:
A user is offered a set of unmarked buttons to push to try to escape a
section of the maze. The buttons all look the same and are a part of the
same panel but perform different actions.

Participants:
- Command: interface for action fired when a button is pressed which
  allows us to fire off different actions for a set of buttons on a board
- Concrete commands: different actions that can be triggered
- Button: contains a reference to a command which is fired off
- Panel: the group of all the buttons
- Others: pressButton and press show how a client would interact with the panel

*/

type command interface {
	execute()
}

type button struct {
	cmd command
}

type panel struct {
	buttons [][]button
}

func (b *button) press() {
	b.cmd.execute()
}

func (p *panel) pressButton(x, y int) {
	p.buttons[x][y].press()
}

type maze struct{}

type shrinkRoomCommand struct {
	m *maze
}

func (c *shrinkRoomCommand) execute() {
	// shrink maze walls at visitor location
}

type slimeCommand struct {
	m *maze
}

func (c *slimeCommand) execute() {
	// slime the visitor
}

type revealDoorCommand struct {
	m *maze
}

func (c *revealDoorCommand) execute() {
	// reveal hidden door
}
