# Restaurant API

## Table of contents

- [General info](#general-info)
- [Technologies](#technologies)
- [Setup](#setup)
  - [Tests](#tests)
  - [Usage](#usage)

## General info

This is a restful application to process backend orders from a restaurant.

## Technologies

These are the technologies used to create this project:

- Golang (Go) 1.15
- [Mux](https://github.com/gorilla/mux)

## Setup

To run this project, build it locally using Go:

```
$ go build -o restaurant -i ./cmd/server
```

```
$ ./restaurant
```

### Tests

In order to run tests:

```
$ go test -v ./...
```

### Usage

Endpoints can be found in `pkg/handler/handler.go`
