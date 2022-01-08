package internal

// Job is type struct will define req text
type Job struct {
	ID   string
	Text string
}

type JobResult struct {
	Frequency map[string]int // word with it's occurance
}
