# aws-health-check

A very simple HTTP server used for health checks from AWS Global Accelerator. This is used to check the "health" of any UDP listener, since the accelerator doesn't support UDP health checks.

# usage

1) Specify the endpoint parameters for the server using the `-port` , `-address` and `path` flags. Default is `0.0.0.0:27015/health`.
2) Point the global accelerator health check endpoint group to the server's configured endpoint.
