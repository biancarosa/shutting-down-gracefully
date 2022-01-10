# Shutting down a Go application gracefully

Ensures all goroutines are finished with their worked or wrapped-up their current work for later.

Code written after a real-life challenge at work on Red Hat's edge-api, and put onto a talk format for DevConf.cz.

## Building the container

`podman build . -t devconf`

## Running the container

`podman run devconf`

## Running on detached mode


`podman run --rm --name devconf -d devconf`

## Stopping the container

`podman stop devconf`

`podman logs -l`