# Health Probe

Health Probe pattern focuses on how an application can communicate its health status to Kubernetes.

## Problem

Kubernetes regularly checks the container process and restarts it if issues are detected. However, just checking the process status is not sufficient, i.e, an application may hang but still be in running state. Kubernetes requires a reliable way to check the health.

## Solution

### 1. Liveness Probe

| Languages |
|:-:|
| [YAML](./livenessProbe.yaml) |
| [Golang](./livenessProbe.go) |

### 2. Readiness Probe

| Languages |
|:-:|
| [YAML](./readinessProbe.yaml) |
| [Golang](./readinessProbe.go) |
