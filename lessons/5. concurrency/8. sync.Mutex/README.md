## sync.Mutex

チャネルが、goroutine間で通信するための素晴らしい方法であることを見てきました。

しかし、通信が必要ない場合はどうでしょう？ コンフリクトを避けるために、一度に1つのgoroutineだけが変数にアクセスできるようにしたい場合はどうでしょうか？

このコンセプトは、[排他制御( *mutual exclusion* )](https://ja.wikipedia.org/wiki/%E6%8E%92%E4%BB%96%E5%88%B6%E5%BE%A1)と呼ばれ、このデータ構造を指す一般的な名前は [**mutex** (ミューテックス)](https://ja.wikipedia.org/wiki/%E3%83%9F%E3%83%A5%E3%83%BC%E3%83%86%E3%83%83%E3%82%AF%E3%82%B9)です。

Goの標準ライブラリは、排他制御を[`sync.Mutex`](https://golang.org/pkg/sync/#Mutex)と次の二つのメソッドで提供します:

- `Lock`
- `Unlock`

`Inc` メソッドにあるように、 `Lock` と `Unlock` で囲むことで排他制御で実行するコードを定義できます。

`Value` メソッドのように、mutexがUnlockされることを保証するために `defer` を使うこともできます。