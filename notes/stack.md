# Stack

a stack can be implemented with a `slice` in go

Push `v` onto stack
```go
stack = append(stack, v)
```

Top of the stack
```go
top := stack[len(stack)-1]
```

Pop the top of the stack off
```go
stack = stack[:len(stack)-1]
```