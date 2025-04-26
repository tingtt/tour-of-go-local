## Creating a slice with make

スライスは、組み込みの `make` 関数を使用して作成することができます。 これは、動的サイズの配列を作成する方法です。

`make` 関数はゼロ化された配列を割り当て、その配列を指すスライスを返します。

```
a := make([]int, 5)  // len(a)=5
```

`make` の3番目の引数に、スライスの容量( *capacity* )を指定できます。 `cap(b)` で、スライスの容量を返します:

```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```