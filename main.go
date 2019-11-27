package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	tasker "github.com/Herlitzd/tasker/lib/core"
	executors "github.com/Herlitzd/tasker/lib/executors"
)

func getConfig(path string) map[string]*tasker.Pipeline {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	byteJSON, _ := ioutil.ReadAll(jsonFile)
	var config tasker.Config

	json.Unmarshal(byteJSON, &config)

	runConfig := buildRunConfig(&config)

	return runConfig
}

func getExecutor(value *string) tasker.Executor {
	return &executors.Local{}
}

func buildRunConfig(config *tasker.Config) map[string]*tasker.Pipeline {

	runConfig := make(map[string]*tasker.Pipeline)

	for pipelineName, pipelineMap := range config.Pipelines {
		pipelineSteps := make(map[string]*tasker.Pipeline)

		executor := getExecutor(&pipelineMap.Executor)

		// Create pipelinesSteps for each known step
		for _, stepStruct := range pipelineMap.Steps {
			task := config.Tasks[stepStruct.Task]
			pipelineSteps[stepStruct.Name] = &tasker.Pipeline{Task: &task, Executor: executor, ForceSuccess: stepStruct.ForceSuccess}
		}

		// Link the pipelinesSteps together
		for _, stepStruct := range pipelineMap.Steps {
			pipelineSteps[stepStruct.Name].OnSuccess = pipelineSteps[stepStruct.OnSuccess]
			pipelineSteps[stepStruct.Name].OnFail = pipelineSteps[stepStruct.OnFail]
		}

		runConfig[pipelineName] = pipelineSteps[pipelineMap.Start]
	}
	return runConfig
}

func main() {
	argsWithProg := os.Args[1:]
	config := getConfig(argsWithProg[0])
	p := config[argsWithProg[1]]
	p.Begin()
	fmt.Print(p.String())
}
