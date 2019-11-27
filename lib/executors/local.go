package executors

import (
	"os/exec"

	tasker "github.com/Herlitzd/tasker/lib/core"
)

type Local struct {
}

func (l *Local) Execute(t *tasker.Task) *tasker.TaskResult {
	command := exec.Command(t.Program, t.Args...)

	stdoutStderr, err := command.CombinedOutput()

	output := string(stdoutStderr)
	result := tasker.TaskResult{BaseTask: t, Output: &output}

	if err != nil {
		result.IsSuccess = false
	} else {
		result.IsSuccess = true
	}

	return &result
}
