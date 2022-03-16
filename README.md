[![Release](https://github.com/looterz/aws-health-check/actions/workflows/release.yml/badge.svg?branch=main)](https://github.com/looterz/aws-health-check/releases)

# aws-health-check

A very simple HTTP server for replying to health checks from AWS Global Accelerator, Application Load Balancer or Network Load Balancer services. This is typically used to check the "health" of any UDP listener, since the services don't support UDP health checks, although it could be used to respond to any health checks in general.

# usage

1) Specify the endpoint parameters for the server using the `-address`, `-port`,  and `-path` flags.
2) Point the global accelerator health check endpoint group to the server's configured endpoint.
