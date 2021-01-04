package pattern

/*

Pattern:
strategy

Premise:
In a maze competition, once all participants reach the end a winner is chosen.
There are different ways of determining a winner, and each way requires
a different way of tallying scores.


Participants:
- Strategy: interface for the family of methods you want to be able to
  switch out other strategies for
- Concrete strategies: different ways the goal of the strategy an be
  accomplished, in this case to determine the maze winner
- Client: the maze, in this case, is the client that

*/

type pickWinnerStrategy interface {
	pickWinner(visitors []*visitor) *visitor
}

type maze struct {
	visitors []*visitor
	strategy pickWinnerStrategy
}

type visitor struct {
	name            string
	flagsCollected  int
	secondsToEscape int
	puzzlesSolved   int
}

func (m *maze) pickWinner() *visitor {
	return m.strategy.pickWinner(m.visitors)
}

type flagsGetWinnerStrategy struct{}

func (s *flagsGetWinnerStrategy) pickWinner(visitors []*visitor) *visitor {
	var winner *visitor

	var max int
	for _, v := range visitors {
		if v.flagsCollected > max {
			max = v.flagsCollected
			winner = v
		}
	}

	return winner
}

type fastEscapeGetWinnerStrategy struct{}

func (s *fastEscapeGetWinnerStrategy) pickWinner(visitors []*visitor) *visitor {
	var winner *visitor

	var min int
	for _, v := range visitors {
		if v.secondsToEscape < min || min == 0 {
			min = v.secondsToEscape
			winner = v
		}
	}

	return winner
}

type puzzlesGetWinnerStrategy struct{}

func (s *puzzlesGetWinnerStrategy) pickWinner(visitors []*visitor) *visitor {
	var winner *visitor

	var max int
	for _, v := range visitors {
		if v.puzzlesSolved > max {
			max = v.puzzlesSolved
			winner = v
		}
	}

	return winner
}
