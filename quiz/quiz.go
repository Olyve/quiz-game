package quiz

import "strings"

type Problem struct {
	question string
	answer   string
}

type Quiz struct {
	problems []Problem
}

func (q Quiz) create(content []byte) {
	// Convert the raw data to a string
	raw_problems := strings.Split(string(content), "\n")

	// Create an array of Problems
	for _, raw := range raw_problems {
		split := strings.Split(raw, ",")
		problem := Problem{
			question: split[0],
			answer:   split[1],
		}
		q.problems = append(q.problems, problem)
	}
}
