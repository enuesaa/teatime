# designdoc
## refactor plan
- パーソナルユースなダッシュボード
- 参照系がメインのユースケースだが、アクションを呼び出せるものもある
- mongodb にデータを保存する
- プラグインでアクションを定義できる
- アクションはデータをフェッチして mongodb に入れるのが主なユースケース
- プラグインの情報は mongodb に入れる

## Create go.work
```bash
go work init . ./plugins/teatime-plugin-coredata
```
