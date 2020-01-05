# ci-info

Get details about the current Continuous Integration environment.

## Installation

```bash
go get -u https://github.com/KlotzAndrew/ci-info

ci-info help
```


## Usage in CLI project

```bash
ci-info isci
# true

ci-info ispr
# flase

ci-info ciname
# Travis CI
```

## Usage in a GO project

```go
package main

import (
	"fmt"

	"github.com/klotzandrew/ci-info/checker"
)

func main() {
  fmt.Printf(
		"ci: %v, pr %v, name: %v",
		checker.IsCI(),
		checker.IsPR(),
		checker.CIName(),
	)
}
```
