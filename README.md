[![Release](https://github.com/looterz/aws-health-check/actions/workflows/release.yml/badge.svg?branch=main)](https://github.com/looterz/aws-health-check/actions/workflows/release.yml)

# aws-health-check

A very simple HTTP server used for health checks from AWS Global Accelerator. This is used to check the "health" of any UDP listener, since the accelerator doesn't support UDP health checks.

# usage

1) Specify the endpoint parameters for the server using the `-port` , `-address` and `path` flags.
2) Point the global accelerator health check endpoint group to the server's configured endpoint.
