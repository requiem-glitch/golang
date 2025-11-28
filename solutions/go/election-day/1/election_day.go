package electionday

import "fmt"
// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var pInitialVotes *int
    pInitialVotes = &initialVotes
    return pInitialVotes
    // panic("Please implement the NewVoteCounter() function")
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
        return 0
    }
    cntr := *counter
    return cntr
    // panic("Please implement the VoteCount() function")
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
    // panic("Please implement the IncrementVoteCount() function")
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	// type ElectionResult struct {
        // Name string
        // Votes int
    // }
    var res *ElectionResult
    res = &ElectionResult{Name: candidateName, Votes: votes}
    return res
    // panic("Please implement the NewElectionResult() function")
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
    return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
    // panic("Please implement the DisplayResult() function")
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate]--
    // panic("Please implement the DecrementVotesOfCandidate() function")
}
