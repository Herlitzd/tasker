package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Task struct {
	Name    string
	Args    string
	Program string
	Script  string
}

type TaskResult struct {
	Output   *string
	Error    *string
	BaseTask *Task
}

func (t Task) String() string {
	return fmt.Sprintf("%v:\n> %v %v\n%v\n", t.Name, t.Program, t.Args, t.Script)
}

func (t TaskResult) String() string {
	empty := ""
	if t.Error == nil {
		t.Error = &empty
	}
	if t.Output == nil {
		t.Error = &empty
	}

	return fmt.Sprintf("Task:\nName:\n%v\nOutput:\n%s\n%s", t.BaseTask, *t.Output, *t.Error)
}

func (t Task) Execute() *TaskResult {
	command := exec.Command(t.Program)
	if t.Args != "" {
		command.Args = []string{t.Args}
	}

	// command.Stdin = strings.NewReader(t.Script)
	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()

	output := out.String()
	result := TaskResult{BaseTask: &t, Output: &output}

	if err != nil {
		errorString := err.Error()
		result.Error = &errorString
	}
	return &result
}
