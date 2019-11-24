package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetConfig(path string) map[string]*Pipeline {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	byteJson, _ := ioutil.ReadAll(jsonFile)
	var config Config

	json.Unmarshal(byteJson, &config)

	runConfig := BuildRunConfig(&config)

	return runConfig
}

func BuildRunConfig(config *Config) map[string]*Pipeline {

	runConfig := make(map[string]*Pipeline)

	for pipelineName, pipelineMap := range config.Pipelines {
		pipelineSteps := make(map[string]*Pipeline)

		for _, stepStruct := range pipelineMap.Steps {
			task := config.Tasks[stepStruct.Task]
			pipelineSteps[stepStruct.Name] = &Pipeline{Task: &task, ForceSuccess: stepStruct.ForceSuccess}
		}

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
	config := GetConfig(argsWithProg[0])
	p := config[argsWithProg[1]]
	p.Begin()
	fmt.Print(p.String())
}
