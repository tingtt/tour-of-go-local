## Exercise: Stringers

`IPAddr` 型を実装してみましょう IPアドレスをドットで4つに区切った( *dotted quad* )表現で出力するため、 `fmt.Stringer` インタフェースを実装してください。

例えば、 `IPAddr{1, 2, 3, 4}` は、 `"1.2.3.4"` として出力するようにします。