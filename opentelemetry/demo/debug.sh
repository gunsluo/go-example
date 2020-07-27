#!/bin/sh
OS=$(uname -s)

export TRACE_ENABLED=true
export METRIC_ENABLED=true
export OTLP_AGENT_EDNPOINT=127.0.0.1:55680


# Postgres
go run main.go all --verbose

