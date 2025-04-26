## Slice literals

スライスのリテラルは長さのない配列リテラルのようなものです。

これは配列リテラルです:

```
[3]bool{true, true, false}
```

そして、これは上記と同様の配列を作成し、それを参照するスライスを作成します:

```
[]bool{true, true, false}
```