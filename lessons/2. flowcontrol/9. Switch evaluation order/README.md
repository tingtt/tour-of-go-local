## Switch evaluation order

Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

(For example,

```
switch i {
case 0:
case f():
}
```

does not call `f` if `i==0`.)

**Note:** Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader.