// cmd/unittestingdev/main.go
package main

import (
"flag"
"log"
"os"

"unittestingdev/internal/unittestingdev"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := unittestingdev.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
