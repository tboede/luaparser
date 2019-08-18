

# Description

luaparser is a go library for parsing go data types into or from lua data types.

# Marshalling

```go
package main

import (
    "fmt"
    lua "github.com/tboede/luaparser"
)

func main() {
    v := true
    s,err := lua.Marshal(v)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(s)
}
```

```go
package main

import (
    "fmt"
    lua "github.com/tboede/luaparser"
)

func main() {
    m := map[string]interface{}
    m["aa"]=true

    s,err := lua.Marshal(m)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(s)
}
```


# Unmarshalling
```go
package main

import (
    "fmt"
    lua "github.com/tboede/luaparser"
)

func main() {
    m := map[string]interface{}
    s := `{2,3,["dd-dd"]=34,gh="aa"}`

    err := lua.Unmarshal([]byte(s),&m)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(m["1"])
    fmt.Println(m["2"])
    fmt.Println(m["dd-dd"])
    fmt.Println(m["gh"])
}
```

```go
package main

import (
    "fmt"
    lua "github.com/tboede/luaparser"
)

func main() {
    var v bool
    s := `false`

    err := lua.Unmarshal([]byte(s),&v)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(v)
}
```
