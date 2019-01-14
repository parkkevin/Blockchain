package work_queue

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	// TODO: initialize struct; start nWorkers workers as goroutines
	q.Jobs = make(chan Worker, maxJobs)
	q.Results = make(chan interface{}, maxJobs)
	for i := uint(0); i < nWorkers; i++ {
		go q.worker()
	}
	return q
}

func (queue WorkQueue) worker() {
	for job := range queue.Jobs {
		queue.Results <- job.Run()
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	close(queue.Jobs)
	for _ = range queue.Jobs {}
}
