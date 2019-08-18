

# Description

luaparser is go library to parse go data types into or from lua data types.

# Marshalling

```go
v := true
s,err := lua.Marshal(v)
```

```go
v := map[string]interface{}
v["1"]=34
v["tree"]="leaf"

s,err := lua.Marshal(v)
```


# Unmarshalling
```go
s := `{true,false,1,2,tree="leaf",[15]="car"}`

m := map[string]interface{}

err := lua.Unmarshal([]byte(s),m)

```

```go
s := `23.45`

var v float64

err := lua.Unmarshal([]byte(s),&v)

```
