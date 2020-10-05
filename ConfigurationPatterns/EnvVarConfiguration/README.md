# EnvVar Configuration

Simplest way to configure applications. Adding data into environment variables.

## Problem

Every non-trivial application requires some configuration for tasks such as accessing data sources. Hardcoding these values into the application code is a bad thing. We externalize this so that values can be changed even after application has been built.

## Solution

| Languages |
|:-:|
| [YAML](./envVarConfiguration.yaml) |
| [Golang](./envVarConfiguration.go) |
