package main

import (
	periodicjob "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/PeriodicJob"
	sidecar "github.com/sharadbhat/KubernetesPatterns/StructuralPatterns/Sidecar"
)

func main() {
	// Behavioral Patterns
	periodicjob.GeneratePeriodicJobSpec()

	// Structural Patterns
	sidecar.GenerateSidecarSpec()
}
