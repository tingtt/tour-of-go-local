## Methods and pointer indirection

下の2つの呼び出しを比べると、ポインタを引数に取る `ScaleFunc` 関数は、ポインタを渡す必要があることに気がつくでしょう:

```
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

メソッドがポインタレシーバである場合、呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができます:

```
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

`v.Scale(5)` のステートメントでは、 `v` は変数であり、ポインタではありません。 メソッドでポインタレシーバが自動的に呼びだされます。 `Scale` メソッドはポインタレシーバを持つ場合、Goは利便性のため、 `v.Scale(5)` のステートメントを `(&v).Scale(5)` として解釈します。