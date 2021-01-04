package pattern

import (
	"time"
)

/*

Pattern:
template method

Premise:
Doors in the maze can take a few formats. To keep things familiar to visitors,
the same basic steps for opening them are always present. However,
different door types are possible whose details extend past the basic steps.

Participants:
- Template: interface for basic steps to follow to open a maze door
- Concrete implementations: implementations of the template which extend it
  in some way specific to the type of entryway

*/

type mazeEntry struct {
	tmpl        entrywayTemplate
	waitTimeout time.Duration
}

func newEntry(
	entry entrywayTemplate,
	waitTimeout time.Duration,
) *mazeEntry {
	var m mazeEntry
	m.tmpl = entry
	m.waitTimeout = waitTimeout
	if waitTimeout == 0 {
		m.waitTimeout = 60 * time.Second
	}

	return &m
}

func (m *mazeEntry) activateEntryway() {
	err := m.tmpl.verifyPasscode()

	if nil != err {
		m.tmpl.displayError(err)
		return
	}

	if err = m.tmpl.open(); nil != err {
		m.tmpl.displayError(err)
	}

	m.waitForPass()

	m.tmpl.close()
}

func (m *mazeEntry) waitForPass() {
	c := make(chan bool)

	go m.tmpl.detectPassed(c)

	select {
	case <-c:
		// pass detected

	case <-time.After(m.waitTimeout):
		// wait time limit passed, allow the door to close
	}
}

type entrywayTemplate interface {
	verifyPasscode() error
	displayError(err error)
	open() error
	detectPassed(c chan bool)
	close()
}

type swingDoor struct {
	code string
}

type swingEntryway struct {
	d *swingDoor
}

func (d *swingEntryway) verifyPasscode() error {
	// verify that the code matches up with a stored code
	// verify the code hasn't already been used with this door
	return nil
}
func (d *swingEntryway) displayError(err error) {
	// display error on the keycard panel
}
func (d *swingEntryway) open() error {
	// activate swing forward mechanism
	return nil
}
func (d *swingEntryway) detectPassed(c chan bool) {
	// check that the weight sensor behind the door was triggered
	time.Sleep(5 * time.Second)

	detected := true
	c <- detected
	return
}
func (d *swingEntryway) close() {
	// activate swing back mechanism
	return
}

type trapdoor struct {
	tangramImg []byte
}

type trapdoorEntryway struct {
	d *trapdoor
}

func (d *trapdoorEntryway) verifyPasscode() error {
	// scan in the image of the tangram
	// verify it's correct
	return nil
}
func (d *trapdoorEntryway) displayError(err error) {
	// read the error out via text to speech and a speaker
	return
}
func (d *trapdoorEntryway) open() error {
	// undo lock keeping door up
	return nil
}
func (d *trapdoorEntryway) detectPassed(c chan bool) {
	// detect thud
	time.Sleep(5 * time.Second)

	detected := true
	c <- detected
	return
}
func (d *trapdoorEntryway) close() {
	// activate swing up mechanism
	return
}
