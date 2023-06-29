# go-shor
Shor's algorithm to factorize a number N

```go
package main

import (
	"fmt"

	"github.com/gophre/shor"
)

var N int = 15

func main() {
	fmt.Println(shor.Factorize(N))
}
```
