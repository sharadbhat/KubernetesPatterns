package main

import (
	// Behavioral Patterns
	batchjob "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/BatchJob"
	periodicjob "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/PeriodicJob"

	// Configuration Patterns
	envvarconfiguration "github.com/sharadbhat/KubernetesPatterns/ConfigurationPatterns/EnvVarConfiguration"

	// Structural Patterns
	initcontainer "github.com/sharadbhat/KubernetesPatterns/StructuralPatterns/InitContainer"
	sidecar "github.com/sharadbhat/KubernetesPatterns/StructuralPatterns/Sidecar"
)

func main() {
	// Behavioral Patterns
	batchjob.GenerateBatchJobSpec()
	periodicjob.GeneratePeriodicJobSpec()

	// Configuration Patterns
	envvarconfiguration.GenerateEnvVarConfigurationSpec()

	// Structural Patterns
	initcontainer.GenerateInitContainerSpec()
	sidecar.GenerateSidecarSpec()
}
