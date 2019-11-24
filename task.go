package main

import (
	"fmt"
	"os/exec"
)

type Task struct {
	Description string
	Args        []string
	Program     string
}

type TaskResult struct {
	Output    *string
	IsSuccess bool
	BaseTask  *Task
}

func (t Task) String() string {

	argString := ""
	for _, arg := range t.Args {
		argString += arg + " "
	}

	return fmt.Sprintf("%v:\n> %v %s\n", t.Description, t.Program, argString)
}

func (t TaskResult) String() string {
	return fmt.Sprintf("Task:\nDescription:\n%v\nSuccess:%v\nOutput:\n%s", t.BaseTask, t.IsSuccess, *t.Output)
}

func (t Task) Execute() *TaskResult {
	command := exec.Command(t.Program, t.Args...)

	stdoutStderr, err := command.CombinedOutput()

	output := string(stdoutStderr)
	result := TaskResult{BaseTask: &t, Output: &output}

	if err != nil {
		result.IsSuccess = false
	} else {
		result.IsSuccess = true
	}

	return &result
}
