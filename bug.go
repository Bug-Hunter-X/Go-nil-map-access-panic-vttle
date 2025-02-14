```go
package main

import "fmt"

func main() {
    var m map[string]int
    m["hello"] = 1 // This will cause a runtime panic if m is nil
    fmt.Println(m["hello"])
}
```