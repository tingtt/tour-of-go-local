## Readers

`io` パッケージは、データストリームを読むことを表現する `io.Reader` インタフェースを規定しています。

Goの標準ライブラリには、ファイル、ネットワーク接続、圧縮、暗号化などで、このインタフェースの [多くの実装](https://golang.org/search?q=Read#Global) があります。

`io.Reader` インタフェースは `Read` メソッドを持ちます:

```
func (T) Read(b []byte) (n int, err error)
```

`Read` は、データを与えられたバイトスライスへ入れ、入れたバイトのサイズとエラーの値を返します。 ストリームの終端は、 `io.EOF` のエラーで返します。

例のコードは、 [`strings.Reader`](//golang.org/pkg/strings/#Reader) を作成し、 8 byte毎に読み出しています。