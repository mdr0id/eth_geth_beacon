#!/bin/bash

docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions
mkdir -p ./data
mkdir -p ./data/grafana
mkdir -p ./data/prometheus
chown -R 2000 ./data/grafana
chown -R 2000 ./data/prometheus