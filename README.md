# Golang Macro Example

Golang does not have built-in macro feature. The closes things it has are **build tags** and **code generators**. But this feature is by nature is standalone - often it operates text rather than specific language, so one may add macroses into Go manually.

This project demonstrates possible implementation and usage of macroses in Go based on **C++** **preprocessor**.

**THIS IS JUST A DEMO**. The code is not production ready - use it just as a hint for your own implementation.

## Requirements

You must have `g++` and `goimports` installed and available in `PATH`.

Run `make build-tools` before running any other command.

## Usage

Use `make` to run demonstrations:

* `make build-tools` - build a code-generation tool once. Required for other commands.
* `make generate` - Generate code with macroses enabled, without running it. The code will be located in folder `generated`.
* `make run` - Generate and run program with macroses and logging enabled.
* `make run-no-logging` - Run program with logging **disabled in compile time**.
* `make run-no-macro` - Run program without using macroses. Logging is evaluated at **runtime**.

## Configuration

You can play with it on your own by changing code or passing following env vars:

* `GO_INCLUDE_DEFINES` - custom comma-separated list of flags to be defined before includes.
  Example:
  `   GO_INCLUDE_DEFINES=DEF_1,DEF2`
  Generates:
  `   #define DEF1`
  `   #define DEF2`
* `GO_INCLUDE_HEADERS_DIR` - path to folder with headers when they are included using relative notion (`"path/to/header.h"` or just `path/to/header.h`). It is not used when header is included using global path (`<path/to/header.h>`).
* `GO_INCLUDE_GEN_DIR` - specify folder for generated files. Defaults to `generated`.

## Advantages of macroses

In original `main.go` you there is following piece of code:

```
LOG("Hello, World! --- %v", toJSON(BigStructure{}))
```

If you run `make generate` you can find next piece of code in `generated/main.go`:

```
if LoggingEnabled {
    fmt.Printf("LOG(macro): "+"Hello, World! --- %v"+"\n", toJSON(BigStructure{}))
}
```

The important changes here are:

* `toJSON(BigStructure{})` won't even execture if `LoggingEnabled` is not set
* Its value won't be passed thought `interface{}` argument

These changes may increase program performance if you have a log of logging.
And this trick can be used not only for `LoggingEnabled` flag, but e.g. for checking currently enabled **logging severity** (as **log4cxx** library does it in `LOG4CXX_*` macroses).

You can even complitely remove parts of code - that's what `make run-no-logging` does - it completelly removed logging from the code.

## Potential use-cases

* Logging: severify check, enabling/disabling, avoiding arguments evaluation.
* Assets: enabling/disabling, avoiding arguments evaluation.
* Error handling (e.g. auto return on error).
* Ternary operator.
* Functional operatators (e.g. `map()`).
* Debugging tools: multi-threading detectors, performance benchmarks.
