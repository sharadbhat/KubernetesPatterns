# Daemon Service

This pattern allows placing and running prioritized, infrastructure-focused Pods on targeted nodes. Used generally by administrators.

## Problem

At an operating system level, a daemon is a long-running, self-recovering computer program that runs as a background process. These generally run in the background and perform tasks such as log collection. Kubernetes provides DaemonSet for this usecase.

## Solution

| Languages |
|:-:|
| [YAML](./daemonService.yaml) |
| [Golang](./daemonService.go) |
