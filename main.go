package main

import (
	// Foundational Patterns
	healthprobe "github.com/sharadbhat/KubernetesPatterns/FoundationalPatterns/HealthProbe"
	managedlifecycle "github.com/sharadbhat/KubernetesPatterns/FoundationalPatterns/ManagedLifecycle"
	predictabledemands "github.com/sharadbhat/KubernetesPatterns/FoundationalPatterns/PredictableDemands"

	// Behavioral Patterns
	batchjob "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/BatchJob"
	daemonservice "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/DaemonService"
	periodicjob "github.com/sharadbhat/KubernetesPatterns/BehavioralPatterns/PeriodicJob"

	// Configuration Patterns
	configurationresource "github.com/sharadbhat/KubernetesPatterns/ConfigurationPatterns/ConfigurationResource"
	envvarconfiguration "github.com/sharadbhat/KubernetesPatterns/ConfigurationPatterns/EnvVarConfiguration"

	// Structural Patterns
	initcontainer "github.com/sharadbhat/KubernetesPatterns/StructuralPatterns/InitContainer"
	sidecar "github.com/sharadbhat/KubernetesPatterns/StructuralPatterns/Sidecar"
)

func main() {
	// Foundational Patterns
	predictabledemands.GenerateResourceLimitSpec()
	healthprobe.GenerateLivenessProbeSpec()
	healthprobe.GenerateReadinessProbeSpec()
	managedlifecycle.GenerateManagedLifecycleSpec()

	// Behavioral Patterns
	batchjob.GenerateBatchJobSpec()
	periodicjob.GeneratePeriodicJobSpec()
	daemonservice.GenerateDaemonServiceSpec()

	// Configuration Patterns
	envvarconfiguration.GenerateEnvVarConfigurationSpec()
	configurationresource.GenerateConfigMapSpec()
	configurationresource.GenerateSecretSpec()

	// Structural Patterns
	initcontainer.GenerateInitContainerSpec()
	sidecar.GenerateSidecarSpec()
}
