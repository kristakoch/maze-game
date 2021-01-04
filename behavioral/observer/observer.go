package pattern

import "fmt"

/*

Pattern:
observer

Premise:
Visitors to the maze can choose to subscribe to hints sporadically
sent out by the maze's help system.

Participants:
- Subscriber: interface for notifying subscribers of events, in this case hints.
  This allows the publisher to notify all subscribers, regardless of how their
  types might differ. identifier added here to enable unsubscribe.
- Concrete publisher: holds references to subscribers and notifies subscribers
  when an event occurs, like a hint being available, so potential users don't
  need to either keep checking back or receive notificactions without a choice.

*/

type subscriber interface {
	update(message string)
	identifier() string
}

type helpSystemPublisher struct {
	subscribers []subscriber
}

type visitorSubscriber struct {
	username string
	hints    []string
}

func (s *visitorSubscriber) update(message string) {
	s.hints = append(s.hints, message)
}

func (s *visitorSubscriber) identifier() string {
	return s.username
}

func (p *helpSystemPublisher) subscribe(s subscriber) {
	p.subscribers = append(p.subscribers, s)
}

func (p *helpSystemPublisher) unsubscribe(identifier string) error {
	for i, s := range p.subscribers {
		if s.identifier() == identifier {
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found by identifier '%s'", identifier)
}

type mazeModeratorSubscriber struct {
	username string
	hints    []string
}

func (s *mazeModeratorSubscriber) update(message string) {
	s.hints = append(s.hints, message)
}

func (s *mazeModeratorSubscriber) identifier() string {
	return s.username
}
