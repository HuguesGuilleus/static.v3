# static.v2

[![PkgGoDev](https://pkg.go.dev/badge/HuguesGuilleus/static.v3)](https://pkg.go.dev/HuguesGuilleus/static.v3)

Create easily a http Handler for static file.

## Installation

```bash
go get -u github.com/HuguesGuilleus/static.v3
```

## Example

```go
package main

import (
	"github.com/HuguesGuilleus/static.v3"
	"log"
	"net/http"
)

func main() {
	// To pass in Dev mode.
	// static.Dev = true

	http.HandleFunc("/", static.Html().File("front/index.html"))
	http.HandleFunc("/style.css", static.Css().File("front/style.css"))
	http.HandleFunc("/app.js", static.Js().File("front/app.js"))
	http.HandleFunc("/favicon.png", static.Png().File("front/favicon.png"))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
```
