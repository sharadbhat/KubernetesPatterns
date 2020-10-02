# Periodic Job

Extends Batch Job by adding a time dimension.

## Problem

To automate system maintenance, administrative tasks or certain specific periodic tasks, we need to use specialized scheduling software or Cron.

A simple service that has to do a certain task daily may end up requiring multiple nodes, a leader, etc. Kubernetes solves this using CronJob.

## Solution

| Languages |
|:-:|
| [YAML](./periodicJob.yaml) |
| [Golang](./periodicJob.go) |
