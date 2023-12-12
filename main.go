package main

import (
	"context"
	"os"

	"github.com/cgradwohl/go-mod-test/components"
)

func main() {
	component := components.Hello("John")
	component.Render(context.Background(), os.Stdout)
}
