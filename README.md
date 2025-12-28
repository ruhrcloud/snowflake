<h1>
  <span>snowflake</span>
</h1>

[![License: MIT](https://img.shields.io/badge/MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/ruhrcloud/snowflake?include_prereleases)](https://github.com/ruhrcloud/snowflake/releases/latest)
![Go Version](https://img.shields.io/github/go-mod/go-version/ruhrcloud/snowflake/main?label=Go)

Snowflakes are distributed and sortable IDs based on Twitter’s Snowflake algorithm.

By design, the IDs generated are intentionally predictable to maintain sortability. Do not use them for sensitive identifiers like session tokens.

![](https://upload.wikimedia.org/wikipedia/commons/5/5a/Snowflake-identifier.png)

## Installation

```bash
go get github.com/ruhrcloud/snowflake
```

## Usage

```go
import (
	"fmt"
	"github.com/ruhrcloud/snowflake"
)

func main() {
    node, _ := snowflake.New(1)
    id, _ := node.Generate()
    fmt.Println(id.String())
}
```

## Documentation

Full documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/ruhrcloud/snowflake).

