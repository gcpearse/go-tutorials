# Tutorial: Get started with Go

Reproduced and modified material from the **Get started with Go** tutorial at [go.dev](https://go.dev/doc/tutorial/getting-started).

## Notes

### Tracking dependencies

Run the following command to enable dependency tracking by creating a `go.mod` file:

`go mod init example/hello`

In this example, `example/hello` is the module path for the new module.

### Running code

Use the following command to run the code:

`go run .`

### Using an external module

Visit **pkg.go.dev** and search for **quote**.

Import the `rsc.io/quote` package and call its `Go()` method.

Run the following command to add the quote module as a requirement and create a `go.sum` file to authenticate the module:

`go mod tidy`

The `rsc.io/quote` module is located and downloaded when `go mod tidy` is run.