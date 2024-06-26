package database

type Participant struct {
	Name    string
	Draw    string
	Contact string
}

type Status struct {
	Participants int
	Bought       int
	Available    int
	Done         bool
	WinnerDraw   int
	WinnerName   string
}

type DrawInfo struct {
	Number    int
	Available bool
}
