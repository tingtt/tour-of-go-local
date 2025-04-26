## Stacking defers

`defer` へ渡した関数が複数ある場合、その呼び出しはスタック( *stack* )されます。 呼び出し元の関数がreturnするとき、 `defer` へ渡した関数は [LIFO(last-in-first-out)](https://ja.wikipedia.org/wiki/LIFO) の順番で実行されます。

`defer` ステートメントについてさらに学ぶには、 [こちら(blog post)](https://blog.golang.org/defer-panic-and-recover)を読んでみてください。