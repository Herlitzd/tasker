package main

type Config struct {
	Tasks     map[string]Task
	Pipelines map[string]PipelineConfig
}

type PipelineConfig struct {
	Start string
	Steps []PipelineStep
}

type PipelineStep struct {
	Name         string
	Task         string
	OnSuccess    string
	OnFail       string
	ForceSuccess bool
}
