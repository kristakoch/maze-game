package pattern

/*

Pattern:
chain of responsibility

Premise:
Multiple methods can handle frightening the person who enters the
maze room based on what the person is most afraid of as detected
by the maze.

Participants:
- Frightener: interface for handling frightens which allows the chaining
  of different frighteners together
- Concrete frighteners: different ways a maze visitor can be frightened,
  can be switched out at will

*/

type frightener interface {
	frighten(afraidOf string)
}

type spiderHandler struct {
	next frightener
	hb   *huntsmanBasket
}

func (s *spiderHandler) frighten(afraidOf string) {
	if afraidOf == "spiders" {
		s.hb.release()
		return
	}

	if nil != s.next {
		s.next.frighten(afraidOf)
	}
}

type huntsmanBasket struct{}

func (hb *huntsmanBasket) release() {
	// release basket of huntsman spiders
}

type heightsHandler struct {
	next frightener
	vr   *headset
}

func (h *heightsHandler) frighten(afraidOf string) {
	if afraidOf == "heights" {
		h.vr.runProgram()
		return
	}

	if nil != h.next {
		h.next.frighten(afraidOf)
	}
}

type headset struct{}

func (h *headset) runProgram() {
	// affix irremovable vr headset with narrow bridge over cliff on windy day
}

type germsHandler struct {
	next frightener
	v    *volunteer
}

func (h *germsHandler) frighten(afraidOf string) {
	if afraidOf == "germs" {
		h.v.dispatch()
		return
	}

	if nil != h.next {
		h.next.frighten(afraidOf)
	}
}

type volunteer struct{}

func (v *volunteer) dispatch() {
	// dispatch volunteer to shake hands while hands are wet and sneeze in face
}
