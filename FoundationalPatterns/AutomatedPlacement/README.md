# Automated Placement

Core function of the Kubernetes scheduler for assigning new Pods to nodes satisfying resource requests and honoring scheduling policies. This pattern describes how placement decisions can be influenced from the outside.

## Problem

Containers have dependancy among themselves, on nodes, resource demands and these vary over time. The cluster itself can undergo changes. The way we place containers impacts the availability, performance and capacity of distributed systems.

## Solution

### 1. Node Affinity

| Languages |
|:-:|
| [YAML](./nodeAffinity.yaml) |
| [Golang](./nodeAffinity.go) |
