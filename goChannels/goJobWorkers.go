package goChannels

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type (
	CustomJobs struct {
		Question float64 `json:"question"`
		JobID    int     `json:"job_ID"`
	}

	CustomResults struct {
		Question  float64       `json:"question"`
		JobID     int           `json:"job_ID"`
		Solution  float64       `json:"solution"`
		WorkerID  int           `json:"worker_ID"`
		TimeTaken time.Duration `json:"time_taken"`
	}
)

// GetRandomNo gets a random value of sleep everytime
func GetRandomValue(max int) (randNo float64) {
	rand.Seed(time.Now().UnixNano())
	randNo = float64(rand.Intn(max))
	return randNo
}

func RunGoJobWorkersExample() {
	jobs := make(chan CustomJobs, 100)
	results := make(chan CustomResults, 100)

	job_count := 100
	worker_count := 5

	// // Spawn Job workers
	for w := 1; w <= worker_count; w++ {
		workerID := w + rand.Intn(5000) + 1000
		go worker(workerID, jobs, results)
	}

	for j := 1; j <= job_count; j++ {
		var customJob CustomJobs
		customJob.JobID = j + 10000
		customJob.Question = GetRandomValue(100)

		// send jobs to queue
		jobs <- customJob
	}
	close(jobs)

	for j := 1; j <= job_count; j++ {
		result := <-results
		fmt.Println("Problem: ", result.Question, " solved by Worker: ", result.WorkerID, " with solution: ", result.Solution, " in time: ", result.TimeTaken)
	}
}

func worker(WorkerID int, jobs <-chan CustomJobs, results chan<- CustomResults) {
	for job := range jobs {
		var customResult CustomResults
		fmt.Println("Job: ", job.JobID, " picked by Worker: ", WorkerID)

		customResult.WorkerID = WorkerID
		customResult.Question = job.Question
		customResult.JobID = job.JobID
		customResult.Solution, customResult.TimeTaken = ExpensiveComputation(job.Question)

		results <- customResult
	}
}

func ExpensiveComputation(num float64) (result float64, time_taken time.Duration) {
	time.Sleep(time.Duration(GetRandomSleep(7)) * time.Second)
	start := time.Now()
	result = math.Exp(num)
	time_taken = time.Since(start)
	return result, time_taken
}
