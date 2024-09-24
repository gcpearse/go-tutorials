# Tutorial: Create a Go module

Reproduced and modified material from the **Create a Go module** tutorial at [go.dev](https://go.dev/doc/tutorial/getting-started).

## Notes

### Calling code from another module

Navigate into the `hello` directory.

Run the following command to replace `example.com/greetings` with `../greetings` to locate the dependency:

`go mod edit -replace example.com/greetings=../greetings`

The `go.mod` file will now contain a replace directive.

Run the following command to synchronise the `example.com/hello` module's dependencies:

`go mod tidy`

The `go.mod` file will now contain a require directive with a pseudo-version number.

### Testing

Go test files have file names ending with `_test.go`.

Navigate into the `greetings` directory.

Run one of the following commands to execute tests in test files:

`go test`
`go test -v`

The `-v` flag provides verbose output.