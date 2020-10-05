# Configuration Resource

Kubernetes provides native configuration resources for general and confidential data, which allows decoupling of configuration lifecycle from the application lifecycle. ConfigMap and Secrets.

## Problem

EnvVar configuration is suitable only for simple configurations. Environment variables can be defined at multiple places which makes it hard to find the definition of a variable. Therefore, it is better to keep all configuration data in a single place. Kubernetes provides this via ConfigMaps and Secrets.

## Solution

### 1. ConfigMap

| Languages |
|:-:|
| [YAML](./configMap.yaml) |
| [Golang](./configMap.go) |

### 1. Secret

| Languages |
|:-:|
| [YAML](./secret.yaml) |
| [Golang](./secret.go) |
