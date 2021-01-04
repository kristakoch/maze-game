package pattern

import "encoding/json"

/*

Pattern:
mediator

Premise:
The host of the maze mediates between the visitors of the maze, the
database of riddle answers, and the engineers that made the game
to simplify interactions.

Participants:
- Mediator: interface for mediated communication within the system
- Concrete mediator: the maze host, which facilitates communication between
  the concrete colleagues within the maze
- Colleages: the different things the host is mediating between

*/

type mediator interface {
	notify(message string)
}

type mazeHostMediator struct {
	visitor  *mazeVisitor
	engineer *mazeEngineer
	db       *mazeDatabase
}

func (m *mazeHostMediator) notify(message string) {
	switch message {
	case "answer submitted":
		ans := m.visitor.getAnswer()
		correct := m.db.checkAnswer(ans)
		if correct {
			m.engineer.openDoor()
		}
	case "fix applied":
		fix := m.engineer.getFix()
		m.db.logFix(fix)
		m.visitor.updateMap(fix)
	}
}

type mazeVisitor struct {
	m           mediator
	answer      string
	mazeDetails json.RawMessage
}

func (v *mazeVisitor) submitAnswer(answer string) {
	v.answer = answer
	v.m.notify("answer submitted")
}

func (v *mazeVisitor) getAnswer() string {
	return v.answer
}

func (v *mazeVisitor) updateMap(fix string) {
	// update the map details based on the fix
}

type mazeDatabase struct {
	m       mediator
	logs    []string
	answers []string
}

func (db *mazeDatabase) checkAnswer(answer string) bool {
	for _, a := range db.answers {
		if answer == a {
			return true
		}
	}
	return false
}

func (db *mazeDatabase) logFix(fix string) {
	db.logs = append(db.logs, fix)
}

type mazeEngineer struct {
	m          mediator
	recentFix  string
	doorIsOpen bool
}

func (e *mazeEngineer) openDoor() {
	e.doorIsOpen = true
}

func (e *mazeEngineer) submitFix(fixName string) {
	e.recentFix = fixName
	e.m.notify("fix applied")
}

func (e *mazeEngineer) getFix() string {
	return e.recentFix
}
