# inverted-index

thread-safe **inverted index** implement by Go language

## Features

## Status

The project is in v1.0.0 version.

## Usage

```go
package main

import (
	"fmt"

	"github.com/cloudfstrife/inverted-index/inverted"
)

var idx inverted.Index

func init() {
	idx = inverted.NewIndex()
}

func main() {
	idx.Push("a", 1)
	idx.Push("a", 2)

	idx.Push("b", 1)
	idx.Push("b", 3)

	idx.Push("c", 5)
	idx.Push("c", 4)
	idx.Push("c", 6)

	fmt.Println(idx.GetAllID("a"))
	fmt.Println(idx.GetAllID("b"))
	fmt.Println(idx.GetAllID("c"))

	fmt.Println("--------------------------------")

	idx.Pop("a", 2)
	idx.Pop("b", 4)
	idx.Pop("c", 6)

	fmt.Println(idx.GetAllID("a"))
	fmt.Println(idx.GetAllID("b"))
	fmt.Println(idx.GetAllID("c"))

}
```

### command

run test

```
make test
```

## Documentation

```
godoc -http=:8090
```

## Contributing

If you are interested in contributing, please checkout [CONTRIBUTING.md](./CONTRIBUTING.md).
We welcome any code or non-code contribution!

## Licensing

licensed by the MIT License. See [LICENSE](./LICENSE) for the full license text.
