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

3. Install templ dependencies into your package.

```bash
go get
```

4. create main.go file

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

4. compile main and include its dependencies

```bash
go build main.go hello_templ.go
```

5. run the executable

```bash
./main
```

6. alternatively you can tell go to compile all go files in the package and then execute main.go

```bash
go run *.go
```

## Create a components package

1. make components directory

```bash
mkdir components
```

2. Only identifiers (like functions, variables, types, etc.) that start with a capital letter are exported from a package and can be accessed from other packages.

Update hello.templ to be capitalizes func and to be part of the components package.

```go
// hello.templ

package components

templ Hello(name string) {
	<div>Hello, { name }</div>
}
```

3. Import the components package in main.

```go
// main.go

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
```

4. Compile main

```bash
go build main.go
```

5. Execute main

```
./main
```
