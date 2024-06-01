# designdoc
## refactor plan
- パーソナルユースなダッシュボード
- 参照系がメインのユースケースだが、アクションを呼び出せるものもある
- without db
- db への保存処理は teapod 側に委譲する. バックエンドは db でもいいし http でもいい. が基本は http を想定する
- terraform ライクにリソースを定義する
- 環境変数 TEATIME_TEAPODS=teatime-links,teatime-something のような形式で値を入れプラグインを有効にする
