## Packages

Goのプログラムは、パッケージ( *package* )で構成されます。

プログラムは `main` パッケージから開始されます。

このプログラムでは `"fmt"` と `"math/rand"` パッケージをインポート( *import* )しています。

規約で、パッケージ名はインポートパスの最後の要素と同じ名前になります。 例えば、インポートパスが `"math/rand"` のパッケージは、 `package rand` ステートメントで始まるファイル群で構成します。

**Note:** ここで実行するプログラムは常に同じ環境で実行されます ので、[擬似乱数](http://ja.wikipedia.org/wiki/%E6%93%AC%E4%BC%BC%E4%B9%B1%E6%95%B0)を返す [`rand.Intn`](https://golang.org/pkg/math/rand/#Rand.Intn) はいつも同じ数を返します。

(数を強制的に変える場合は、乱数生成でシードを与える必要があります。[`rand.Seed`](https://golang.org/pkg/math/rand/#Seed)を見てみてください。 playground 上での時間は一定なので他のものをシードとして使う必要があります。)