# Introduction

This repo contains a solution to the SalesLoft offline exerise (v2).

# Architecture

In deciding how to solve this, I decided to choose two technologies
that were new to me:

- Vue.js for the front-end.
- The Chi project which uses Go for the back-end.

*NB*: It is entirely possible that these choices will change during the course of development. ;)

# Development

This application uses Docker to manage compilation, testing, and locally running the server.

## Build
To build the application, run the `build.sh` script. This will create a container,
run the unit tests with coverage, and conditionally compile the application in the container.

## Run
The applicatioon has two run flags:
- `--apikey` is for the SalesLoft api key.
- `--port` is the port for service. The application's default is `3000`.

To run the application:
- If running locally after compilation (e.g. via `go build ...`), exeucte the application binary (e.g. `slpeople`) with at least the `--apikey` flag.
- To run the application using the container, you can use the `run.sh` script and provide the API Key and port.
-- Using `run.sh`: `./run.sh "$apikey" "$port"
-- This will execute: `> docker run --rm -it -p $port:$port slpeople "$apikey" "$port"
