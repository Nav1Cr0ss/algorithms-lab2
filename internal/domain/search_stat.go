package domain

type SearchStats struct {
	Iterations      int
	VisitedStates   int
	StatesInMemory  int
	MaxQueueSize    int
	MaxDepthReached int
	FoundGoal       bool
}
