package pattern

import "sort"

/*

Pattern:
iterator

Premise:
The maze game has a set of layouts the user can choose. When choosing
a plan, they may want to view the list of possible plans in order of
difficulty or grouped by maze material.

Participants:
- Iterator: interface that allows general iteration over the plans
- Concrete iterators: different ways of listing out all plans
- Collection: interface for collections, which create an iterator
- Concrete collections: different types of collections

Note: using this with a slice is kinda overkill because this is meant to
be used with complex data structures, but using here for simplicity.

*/

type iterator interface {
	next() *plan
}

type layoutCollection interface {
	createIterator(plans []*plan) iterator
}

type plan struct {
	difficulty int
	walls      material
}

type material int

const (
	matHedge material = iota
	matCorn
	matBrick
)

type difficultyIterator struct {
	current int
	plans   []*plan
}

func (i *difficultyIterator) next() *plan {
	if len(i.plans) == 0 {
		return nil
	}

	next := i.plans[i.current+1]
	i.current++

	if i.current == (len(i.plans)) {
		i.current = 0
	}

	return next
}

type difficultyCollection []*plan

func (c *difficultyCollection) createIterator(plans []*plan) iterator {
	sort.Slice(plans, func(i, j int) bool {
		return plans[i].difficulty < plans[j].difficulty
	})

	return &difficultyIterator{
		plans: plans,
	}
}

type materialIterator struct {
	current int
	plans   []*plan
}

func (i *materialIterator) next() *plan {
	if len(i.plans) == 0 {
		return nil
	}

	next := i.plans[i.current]
	i.current++

	if i.current == (len(i.plans)) {
		i.current = 0
	}

	return next
}

type materialCollection []*plan

func (c *materialCollection) createIterator(plans []*plan) iterator {
	sort.Slice(plans, func(i, j int) bool {
		return plans[i].walls < plans[j].walls
	})

	return &materialIterator{
		plans: plans,
	}
}
