package repository

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/ajaymahar/vertisystem/internal"
	"github.com/google/uuid"
)

var (
	ErrJobNotFound error = errors.New("job not found")
)

// #############################################
// TopWord will hold the words with Frequency
type TopNWord struct {
	Key   string
	Value int
}

type TopNWordList []TopNWord

//
// Custom sorting logic
func (tw TopNWordList) Len() int {
	return len(tw)
}

func (tw TopNWordList) Swap(i, j int) {
	tw[i], tw[j] = tw[j], tw[i]
}

func (tw TopNWordList) Less(i, j int) bool {
	return tw[i].Value < tw[j].Value
}

// ############################################

// StubRepository implementation for the JobRepo
type StubRepository struct {
	// temp in memory data store
	jobs map[string]string
	// jobs map[internal.Job]internal.JobResult

	// temp result datastre
	result map[string]internal.JobResult
}

// NewStubRepository is factory function to create new
func NewStubRepository() *StubRepository {
	return &StubRepository{
		jobs:   make(map[string]string),
		result: make(map[string]internal.JobResult),
	}
}

// Create JobRepository implementation
func (sr *StubRepository) Create(ctx context.Context, text string) (internal.Job, error) {
	j := internal.Job{
		ID:   getNewID(),
		Text: text,
	}
	// sr.jobs[j] = internal.JobResult{}
	sr.jobs[j.ID] = j.Text
	// fmt.Printf("%+v", sr.jobs)

	go func() {
		sr.processText(j)
	}()
	return j, nil
}

func (sr *StubRepository) Get(ctx context.Context, id string) (internal.JobResult, error) {
	// result, ok := sr.result[id]
	// if !ok {
	// 	return internal.JobResult{}, ErrJobNotFound
	// }
	result, err := sr.getTopNWords()
	if err != nil {
		return internal.JobResult{}, fmt.Errorf(err.Error())
	}
	return result, nil
}

//Generate new ID for each job
func getNewID() string {
	return uuid.NewString()
}

// processText is local func to process the provided text
// it will be responsable to extract the text
// get all the words from the req text
// count occurance of each word
// and store it to result datastore
func (sr *StubRepository) processText(j internal.Job) {
	words := strings.Split(j.Text, " ")
	// fmt.Println("all words here: ", words)
	jr := internal.NewJobResult()
	fmt.Println("jr:", jr)
	for _, word := range words {
		c := strings.Count(j.Text, word)
		// fmt.Println("word: ", word)
		// fmt.Println("count: ", c)
		jr.Frequency[word] = c
		// fmt.Printf("%+v", jr)
	}
	sr.result[j.ID] = *jr
	// fmt.Printf("%+v", sr.result)
}

func (sr *StubRepository) getTopNWords() (internal.JobResult, error) {
	w := TopNWordList{}
	for _, v := range sr.result {
		w = make(TopNWordList, len(v.Frequency))
		i := 0
		for k, v := range v.Frequency {
			w[i] = TopNWord{
				Key:   k,
				Value: v,
			}
			i++
		}
	}
	// sorting
	sort.Sort(sort.Reverse(w))

	// var jr internal.JobResult
	r := internal.NewJobResult()
	for _, kv := range w {
		fmt.Printf("value: %+v", kv)
		r.Frequency[kv.Key] = kv.Value
	}
	return *r, nil
	// keys := make([]int, )

	// hack := map[int]string{}
	// hackkeys := []int{}
	// for _, v := range sr.result {
	// 	for key, val := range v.Frequency {
	// 		hack[val] = key
	// 		hackkeys = append(hackkeys, val)
	// 	}
	// 	sort.Ints(hackkeys)
	// }
	// sort.Slice(hackkeys, func(i, j int) bool {
	// 	return hackkeys[i] > hackkeys[j]
	// })
	// fmt.Println("hack: ", hack)
	// fmt.Println("hackkeys: ", hackkeys)
}
