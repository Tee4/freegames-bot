package scheduler

type Job interface {
	Run() error
}

type Scheduler struct {
	jobs []Job
}

func New() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Add(job Job) {
	s.jobs = append(s.jobs, job)
}

func (s *Scheduler) Run() error {
	for _, job := range s.jobs {
		if err := job.Run(); err != nil {
			return err
		}
	}
	return nil
}
