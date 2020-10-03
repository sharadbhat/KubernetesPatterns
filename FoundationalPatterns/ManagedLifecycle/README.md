# Managed Lifecycle

Applications have to listen to the events emitted by the managing platform and adapt accordingly. The Managed Lifecycle pattern describes how applications can and should react to these

## Problem

The platform on which the containerized application is running can stop the application at any moment for any reason. It is up to the application to decide how to react to different such events. Kubernetes' PostStart and PreStop hooks allow such an adaptation.

## Solution

| Languages |
|:-:|
| [YAML](./managedLifecycle.yaml) |
| [Golang](./managedLifecycle.go) |
