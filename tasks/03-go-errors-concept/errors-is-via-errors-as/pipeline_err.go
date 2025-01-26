package pipe

import (
	"errors"
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

func IsPipelineError(err error, user, pipelineName string) bool {
	// Реализуй меня.
	// if errors.Is(err, &PipelineError{
	// 	User: user,
	// 	Name: pipelineName,
	// }) {
	// 	return true
	// }
	// var n *PipelineError = &PipelineError{
	// User:        user,
	// Name:        pipelineName,
	// FailedSteps: []string{pipelineName},
	// }

	var n *PipelineError

	switch {
	case errors.As(err, &n):
		if n.User == user && n.Name == pipelineName {
			return true
		}
		return false
	default:
		return false
	}
}
