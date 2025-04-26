## Images

[`image`](https://golang.org/pkg/image/#Image) パッケージは、以下の `Image` インタフェースを定義しています：

```
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

Note: `Bounds` メソッドの戻り値である `Rectangle` は、 `image` パッケージの [`image.Rectangle`](https://golang.org/pkg/image/#Rectangle) に定義があります。

(詳細は、 [このドキュメント](https://golang.org/pkg/image/#Image) を参照してください。)

`color.Color` と `color.Model` は共にインタフェースですが、定義済みの `color.RGBA` と `color.RGBAModel` を使うことで、このインタフェースを無視できます。 これらのインタフェースは、[image/color](https://golang.org/pkg/image/color/) パッケージで定義されています。