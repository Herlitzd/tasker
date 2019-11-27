package core

import (
	"fmt"
)

type Pipeline struct {
	Task         *Task
	TaskResult   *TaskResult
	OnSuccess    *Pipeline
	OnFail       *Pipeline
	ForceSuccess bool
	Executor     Executor
}

func (p Pipeline) String() string {
	if p.TaskResult != nil {
		s := fmt.Sprintf("%v", p.TaskResult)
		if p.OnSuccess != nil {
			s = fmt.Sprintf("%s\nSuccess:\n%s\n", s, p.OnSuccess)
		}
		if p.OnFail != nil {
			s = fmt.Sprintf("%s\nFail:\n%s\n", s, p.OnFail)
		}
		return s
	}
	return ""
}

func (p *Pipeline) Begin() {
	if p.Task != nil {
		p.TaskResult = p.Executor.Execute(p.Task)
		switch {
		case p.TaskResult.IsSuccess, p.ForceSuccess:
			if p.OnSuccess != nil {
				p.OnSuccess.Begin()
			}
		case !p.TaskResult.IsSuccess:
			if p.OnFail != nil {
				p.OnFail.Begin()
			}
		}
	}
}
