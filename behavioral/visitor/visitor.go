package pattern

/*

Plan:
visitor

Premise:
All rooms often need to be modified, and new room types are also being added all
the time. But it's costly to change the rooms themselves, so instead of changing
the thing that generates the rooms engineers made a one-time change to make the
rooms accessible to a robot which can perform specialized tasks based on the room.

Participants:
- Visitor: interface for the type of visitor that can be accepted by a
  place being visited. This allows changing logic in the visitor instead of on
  the thing being visited
- Concrete visitors: perform an operation over all nodes in the system, and
  have a method for each type of node being visited. In this case, adds treasure.

Note: Go does not let you override methods, which is what stops us from being
able to have a generic visit(v visitor) method for all types of visits.

*/

type visitor interface {
	visitHallway(r *roomHallway)
	visitPanicRoom(r *roomPanicRoom)
	visitGrandBallroom(r *roomGrandBallroom)
}

type roomHallway struct{}
type roomPanicRoom struct{}
type roomGrandBallroom struct{}

type addTreasureVisitor struct{}

func (v *addTreasureVisitor) visitHallway(r *roomHallway) {
	// add a low level treasure along the walls
}
func (v *addTreasureVisitor) visitPanicRoom(r *roomPanicRoom) {
	// add a mid level treasure right in the middle
}
func (v *addTreasureVisitor) visitGrandBallroom(r *roomGrandBallroom) {
	// add a large treasure 3/4 of the way into the room
}

func (r *roomHallway) accept(v visitor) {
	v.visitHallway(r)
}
func (r *roomPanicRoom) accept(v visitor) {
	v.visitPanicRoom(r)
}
func (r *roomGrandBallroom) accept(v visitor) {
	v.visitGrandBallroom(r)
}
