package main

import (
	// Behavioral Patterns
	batchjob "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/BatchJob"
	periodicjob "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/PeriodicJob"

	// Structural Patterns
	sidecar "github.com/sharadbhat/KubernetesPatterns/StructuralPatterns/Sidecar"
)

func main() {
	// Behavioral Patterns
	batchjob.GenerateBatchJobSpec()
	periodicjob.GeneratePeriodicJobSpec()

	// Structural Patterns
	sidecar.GenerateSidecarSpec()
}
