package core

/**

Lambda Running in AWS
	Receives Event
		Hits s3 bucket get config (Maybe Git)
		Processes Command:
			Each Task
*/
type Executor interface {
	Execute(task *Task) *TaskResult
}
