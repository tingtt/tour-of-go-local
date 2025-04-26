## Range continued

インデックスや値は、 " `_` "(アンダーバー) へ代入することで捨てることができます。

```
for i, _ := range pow
for _, value := range pow
```

もしインデックスだけが必要なのであれば、2つ目の値を省略します。

```
for i := range pow
```