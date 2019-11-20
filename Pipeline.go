package main

import (
	"fmt"
)

type Pipeline struct {
	Task         *Task
	TaskResult   *TaskResult
	OnSuccess    *Pipeline
	OnFail       *Pipeline
	ForceSuccess bool
}

func (p Pipeline) String() string {
	if p.TaskResult != nil {
		s := fmt.Sprintf("%v", p.TaskResult.String())
		if p.OnSuccess != nil {
			s = fmt.Sprintf("%s\nSuccess:\n%s", s, p.OnSuccess.String())
		}
		if p.OnFail != nil {
			s = fmt.Sprintf("%s\nFail:\n%s", s, p.OnFail.String())
		}
		return s
	}
	return ""
}

func (p *Pipeline) Begin() {
	if p.Task != nil {
		p.TaskResult = p.Task.Execute()
		switch {
		case p.TaskResult.Error == nil:
			fallthrough
		case p.ForceSuccess:
			if p.OnSuccess != nil {
				p.OnSuccess.Begin()
			}
		case p.TaskResult.Error != nil:
			if p.OnFail != nil {
				p.OnFail.Begin()
			}
		}
	}
}
