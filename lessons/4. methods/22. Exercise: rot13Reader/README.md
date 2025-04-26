## Exercise: rot13Reader

よくあるパターンは、別の `io.Reader` をラップし、ストリームの内容を何らかの方法で変換する[io.Reader](https://golang.org/pkg/io/#Reader)です。

例えば、 [gzip.NewReader](https://golang.org/pkg/compress/gzip/#NewReader) は、 `io.Reader` (gzipされたデータストリーム)を引数で受け取り、 `*gzip.Reader` を返します。 その `*gzip.Reader` は、 `io.Reader` (展開したデータストリーム)を実装しています。

`io.Reader` を実装し、 `io.Reader` で[ROT13](https://ja.wikipedia.org/wiki/ROT13) 換字式暗号( *substitution cipher* )をすべてのアルファベットの文字に適用して読み出すように `rot13Reader` を実装してみてください。

`rot13Reader` 型は提供済みです。 この `Read` メソッドを実装することで `io.Reader` インタフェースを満たしてください。