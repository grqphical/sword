# sword - A lightweight web framework built around the Golang Standard Library HTTP server
[![Go Tests](https://github.com/grqphical/sword/actions/workflows/go.yml/badge.svg)](https://github.com/grqphical/sword/actions/workflows/go.yml)

I made this as a personal tool to improve the design of my APIs. The only difference this library makes is that it allows for better management of middleware and explicitly returning errors from handlers

## Installation

```bash
$ go get github.com/grqphical/sword
```

## Basic Usage

Here is a basic example of creating a web server with Sword:

```go
package main

import (
    "github.com/grqphical/sword"
    "fmt"
    "os"
    "errors"
    "log"
)

// by default Sword returns all errors as JSON but you can override that behaviour by creating
// a custom error handler.
func customErrorHandler(w http.ResponseWriter, err error) {
    var swordError sword.HandlerError
    if errors.As(err, &swordError) {
        w.WriteHeader(swordError.code)
    }

    fmt.Fprintln(os.Stderr, "ERROR: %s\n", err)
}

func main() {
    // you can provide a configuration or leave it as nil to use the defaults:
    // address = ":5000"
    // errorHandler = defaultErrorHandler
    r := sword.NewRouter(&sword.RouterConfig{
        address: ":8000",
        errorHandler: customErrorHandler
    })

    // Sword works with Golang's built-in routing so things such as methods and wildcards are allowed
    r.RouteFunc("GET /", func(w http.ResponseWriter, r *http.Request) error {
		w.Write([]byte("Hello, World!"))
		return nil
	})

    // with Sword you explicitly return errors whenever an error occurs in your handlers
    r.RouteFunc("GET /error", func(w http.ResponseWriter, r *http.Request) error {
		return sword.Error(http.StatusInternalServerError, "error")
	})

    log.Fatal(r.ListenAndServe())
}
```

## License

Sword is licensed under the MIT License
