## Exercise: Equivalent Binary Trees

同じ数列を保持するような、形の異なる[二分木( *binary tree* )](https://ja.wikipedia.org/wiki/%E4%BA%8C%E5%88%86%E6%9C%A8)は多く存在し得ます。 例えば、ここに数列 1, 1, 2, 3, 5, 8, 13 を保持する2つの二分木があります。

![](/content/img/tree.png)

2つの二分木が同じ数列を保持しているかどうかを確認する機能は、多くの言語においてかなり複雑です。

シンプルな解決方法を記述するために、Goの並行性( *concurrency* )とチャネルを利用してみます。

例では、型を以下のように定義している `tree` パッケージを利用します:

```
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```

[続く...](javascript:click%28'.next-page'%29)