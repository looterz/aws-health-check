# aws-health-check

A very simple HTTP server used for health checks from AWS Global Accelerator.

# usage

1) Specify the port to listen on for incoming HTTP health checks with the `-port` flag.
2) Point the global accelerator health check endpoint group to the server's `/health` endpoint.
