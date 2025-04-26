## Methods and pointer indirection (2)

逆にも見てみましょう。

変数の引数を取る関数は、特定の型の変数を取る必要があります:

```
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

メソッドが変数レシーバである場合、呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができます:

```
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

この場合、 `p.Abs()` は `(*p).Abs()` として解釈されます。