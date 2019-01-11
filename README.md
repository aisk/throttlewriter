# throttlewriter

[![GoDoc](https://godoc.org/github.com/aisk/throttlewriter?status.svg)](https://godoc.org/github.com/aisk/throttlewriter)

![ThrottlePicture](https://media.defense.gov/2006/Jul/26/2000547457/-1/-1/0/060726-F-1234P-021.JPG)

Simple and naïve throttle writer for go.

## Example

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aisk/throttlewriter"
)

func main() {
	writer := throttlewriter.New(os.Stdout, 10, time.Second)
	i := 0
	for {
		i++
		fmt.Fprintln(writer, i)
	}
}
```

```sh
$ go run example.go
1
2
3
4
5
4383562
4383563
8795986
8795987
13150686
13150687
^Csignal: interrupt
```

## License

Licensed under either of
 * MIT license ([LICENSE](LICENSE) or http://opensource.org/licenses/MIT)

at your option.
