## Prerequisites

- install Go
- - the default directory to downloaded go packages is `$HOME/go/bin`

- install templ: `go install github.com/a-h/templ/cmd/templ@latest`
- - may need to update PATH to include the default GOBIN. add `export PATH=$PATH:$HOME/go/bin` to the end of your .zshrc

## Create a new Go Module

Create a New Directory for Your Project:

```bash
mkdir myproject
cd myproject
```

Initialize the Module:
This command creates a go.mod file in your directory.

```bash
go mod init github.com/yourusername/myproject
```

## Getting started with templ

1. create templ component

```go
// hello.templ

package main

templ hello(name string) {
	<div>Hello, { name }</div>
}
```

2. compile the component to go

```bash
templ generate
```

3. create main.go file

```go
// main.go

package main

import (
	"context"
	"os"
)

func main() {
  // 'hello' is the templ component we made in hello.templ
	component := hello("John")
	component.Render(context.Background(), os.Stdout)
}
```

3. compile main and include its dependencies

```go
go build main.go hello_templ.go
```

4. run the executable

```bash
./main
```

## Create a components package
