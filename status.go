package status

type Status string

const (
	OK          Status = "OK"
	DOWN        Status = "DOWN"
	MAINTENANCE Status = "MAINTENANCE"
)

type Checker interface {
	Name() string
	Status() Status
}

type Pool interface {
	Status() Status
	Details() map[string]Status
}

type checkResult struct {
	name   string
	status Status
}

type CheckerPool struct {
	checkers []Checker
}

func NewCheckerPool(checkers ...Checker) *CheckerPool {
	return &CheckerPool{checkers}
}

func (cp *CheckerPool) Status() Status {
	resultChannel := cp.check()

	for i := 0; i < len(cp.checkers); i++ {
		result := <-resultChannel
		if result.status != OK {
			return result.status
		}
	}

	return OK
}

func (cp *CheckerPool) Details() map[string]Status {
	details := make(map[string]Status)
	resultChannel := cp.check()

	for i := 0; i < len(cp.checkers); i++ {
		result := <-resultChannel

		details[result.name] = result.status
	}

	return details
}

func (cp *CheckerPool) check() <-chan checkResult {
	resultChannel := make(chan checkResult, len(cp.checkers))

	for _, checker := range cp.checkers {
		go func(ch Checker) {
			resultChannel <- checkResult{name: ch.Name(), status: ch.Status()}
		}(checker)
	}

	return resultChannel
}
