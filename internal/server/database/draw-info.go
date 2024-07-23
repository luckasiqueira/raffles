package database

type Participant struct {
	Name    string
	Draw    string
	Contact string
	Paid    int8
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
