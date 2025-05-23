## Goroutines

A *goroutine* is a lightweight thread managed by the Go runtime.

```
go f(x, y, z)
```

starts a new goroutine running

```
f(x, y, z)
```

The evaluation of `f`, `x`, `y`, and `z` happens in the current goroutine and the execution of `f` happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The [`sync`](/pkg/sync/) package provides useful primitives, although you won't need them much in Go as there are other primitives. (See the next slide.)