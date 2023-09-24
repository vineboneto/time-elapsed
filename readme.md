# Time Elapsed

Package elapsed provides a simple utility for measuring the elapsed time of code execution in Go.

The TimeElapsed structure is used to track the start and end times of a code block and calculate the elapsed time in milliseconds.

## Installation

You can install the `elapsed` library using the `go get` command:

```bash
go get github.com/vineboneto/time-elapsed
```

Once installed, you can import the library into your Go code by using the import statement like this:

```go
package main

import (
	"github.com/vineboneto/time-elapsed"
	"time"
)

func main() {
	elapsed := timeelapsed.TimeElapsed{}

	elapsed.Start()

	time.Sleep(2 * time.Second)

	elapsed.End("Test1") // 2000ms

	elapsed.Start()

	time.Sleep(2 * time.Second)

	elapsed.End("Test2") // 2000ms
}
```
