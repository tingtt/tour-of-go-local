## Stringers

もっともよく使われているinterfaceの一つに [`fmt` パッケージ](//golang.org/pkg/fmt/) に定義されている [`Stringer`](//golang.org/pkg/fmt/#Stringer) があります:

```
type Stringer interface {
    String() string
}
```

`Stringer` インタフェースは、stringとして表現することができる型です。 `fmt` パッケージ(と、多くのパッケージ)では、変数を文字列で出力するためにこのインタフェースがあることを確認します。