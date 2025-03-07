# Contributing

## Contributor License Agreement
 
Contributions to this project must be accompanied by a **Contributor License Agreement**. 

You or your employer retain the copyright to your contribution: Accepting this agreement gives us permission to use and redistribute your contributions as part of the project.

## Pull Requests

Pull requests must pass all [CI/CD](#cicd) measures.

## CI/CD

### Static Code Analysis

Copygen uses [golangci-lint](https://github.com/golangci/golangci-lint) in order to statically analyze code. You can install golangci-lint with `go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5` and run it using `golangci-lint run`. If you receive a `diff` error, you must add a `diff` tool in your PATH. There is one located in the `Git` bin.

If you receive `File is not ... with -...`, use `golangci-lint run --disable-all --no-config -Egofmt --fix`.