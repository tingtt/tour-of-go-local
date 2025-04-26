## Exercise: Errors

`Sqrt` 関数を [以前の演習](/flowcontrol/8) からコピーし、 `error` の値を返すように修正してみてください。

`Sqrt` は、複素数をサポートしていないので、負の値が与えられたとき、nil以外のエラー値を返す必要があります。

新しい型:

```
type ErrNegativeSqrt float64
```

を作成してください。

そして、 `ErrNegativeSqrt(-2).Error()` で、 `"cannot Sqrt negative number: -2"` を返すような:

```
func (e ErrNegativeSqrt) Error() string
```

メソッドを実装し、 `error` インタフェースを満たすようにします。

**注意:** `Error` メソッドの中で、 `fmt.Sprint(e)` を呼び出すことは、無限ループのプログラムになることでしょう。 最初に `fmt.Sprint(float64(e))` として `e` を変換しておくことで、これを避けることができます。 なぜでしょうか？

負の値が与えられたとき、 `ErrNegativeSqrt` の値を返すように `Sqrt` 関数を修正してみてください。