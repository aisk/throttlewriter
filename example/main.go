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
