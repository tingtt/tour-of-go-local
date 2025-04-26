## Default Selection

どの `case` も準備ができていないのであれば、 `select` の中の `default` が実行されます。

ブロックせずに送受信するなら、 `default` の `case` を使ってください:

```
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```