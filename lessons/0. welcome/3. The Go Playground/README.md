## The Go Playground

Go Tour は、 [golang.org](https://golang.org/) のサーバで稼働しているウェブサービスの [Go Playground](https://play.golang.org/) を利用しています。

このウェブサービスは、Goのコードを受け取ると[サンドボックス](https://ja.wikipedia.org/wiki/%E3%82%B5%E3%83%B3%E3%83%89%E3%83%9C%E3%83%83%E3%82%AF%E3%82%B9_%28%E3%82%BB%E3%82%AD%E3%83%A5%E3%83%AA%E3%83%86%E3%82%A3%29)内でコンパイル、リンク、実行し、 実行結果の出力を返します。

ただし、 playground で実行できるプログラムには制限があります:

- playground 上はいつも "2009-11-10 23:00:00 UTC" (この値の意味は、読者の楽しみのために残しておきます(^^))です。これにより、同じ出力結果を得ることが容易になります。

<!--THE END-->

- 実行時間とCPU、メモリの使用量に制限があり、プログラムは外部のネットワークホストへアクセス出来ません。

playgroundは、最新の安定バージョンのGoを利用します。

詳細は、 "[Inside the Go Playground](https://blog.golang.org/playground)" を読んでみてください。