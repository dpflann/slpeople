# Introduction

This repo contains a solution to the SalesLoft offline exerise (v2).
This repo contains a solution to the SalesLoft offline exerise (v2). This application
 will:
 - List people from the SalesLoft API.
 - Provide character frequency analysis of the characters in the people's email address.
 - Provide an initial solution for identifying duplicate email addresses.
 
 ![application_preview](/static/images/simple.app.view.png)

# Architecture

In deciding how to solve this, I decided to choose two technologies
that were new to me:

- ![Vue.js](https://vuejs.org/) for the front-end.
- The ![Chi](https://github.com/go-chi/chi) project which uses Go for the back-end.

*NB*: It is entirely possible that these choices will change during the course of development. ;)

# Development

This application uses Docker to manage compilation, testing, and locally running the server.

# Requirements and Dependencies
- Go
  - ![core language](https://golang.org/): v11.4
  - ![dep](https://golang.github.io/dep/) (dependency management): v0.5.0
- Vue.js
  - core: v2.5.21
  - ![Vue Resource](https://github.com/pagekit/vue-resource): v1.5.1
- ![d3.js](https://d3js.org/): v3
- ![Docker](https://www.docker.com/): v18.09.0

## Build
To build the application, run the `build.sh` script in the project's root directory. This will create a container,
run the unit tests with coverage, and conditionally compile the application to `slpeople.app` within the container.
- `> ./build.sh`

## Test
Testing is written into the Dockerfile (subject to change, convenient for now), and it uses go's builtin test package.
As stated above, executing the `build.sh` script will run tests before compilation.

To test the application code (mainly go code at this moment), run the following command at the project root:
- `> go test -v -cover ./...`
  - This will execute all tests in all packages.
The output will look like the following:
<pre><code>
?       github.com/slpeople     [no test files]
?       github.com/slpeople/app [no test files]
=== RUN   TestCharacterFrequency
--- PASS: TestCharacterFrequency (0.00s)
=== RUN   TestSortedCharFrequency
--- PASS: TestSortedCharFrequency (0.00s)
=== RUN   TestCharacterFrequencies
--- PASS: TestCharacterFrequencies (0.00s)
PASS
coverage: 62.2% of statements
ok      github.com/slpeople/characters  0.005s  coverage: 62.2% of statements
=== RUN   TestFindPossibleDuplicates
--- PASS: TestFindPossibleDuplicates (0.00s)
    duplicates_test.go:382: len of result: 0
PASS
coverage: 76.0% of statements
ok      github.com/slpeople/duplicates  0.009s  coverage: 76.0% of statements
?       github.com/slpeople/errors      [no test files]
?       github.com/slpeople/salesloftapi        [no test files]
</code></pre>


## Run
The applicatioon has two run flags:
- `--apikey` is for the SalesLoft api key.
- `--port` is the port for service. The application's default is `3000`.

To run the application:
- If running locally after compilation (e.g. via `go build ...`), exeucte the application binary (e.g. `slpeople`) with at least the `--apikey` flag.
- To run the application using the container, you can use the `run.sh` script and provide the API Key and port.
  - Using `run.sh`: `./run.sh "$apikey" "$port"`
  - This will execute: `> docker run --rm -it -p $port:$port slpeople "$apikey" "$port"`

## Test, Build, Run
To do all this at once (test, build, run), run these contingent commands:
- `./build.sh && ./run.sh "$apikey" "$port"`

# TODO
Future work:
- Add more tests for each go package.
- Add more tests for the UI and Vue.js components.
- Add integration tests for network requests -- expected requests and responses.
- Get help from a designer and UI expert to improve the UI.
- Improve performance for listing people, characer frequency analysis, and fuzzy-matching for duplicates.
- There is always more work...
