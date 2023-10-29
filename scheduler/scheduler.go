package schedulerd

type Scheduler interface {
	SelectCandidateNodes()
	Score()
	Pick()
}
