package pipe

import (
	"fmt"
)

type PipelineError struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

// Добавь метод Is для типа *PipelineError.
func (p *PipelineError) Is(target error) bool {
    if p == nil {
		return true
	}

	s, ok := target.(*PipelineError)
    if ok {
		if (p.User == s.User && p.Name == s.Name) {
			return true
		}
	}    
	return false
} 


