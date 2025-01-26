# teatime
My Daily Dashboard. But work in progress..

## Dev Server
```bash
# start mongodb
docker compose up

# create go.work
go work init . ./teapods/teapod-links

# install teapod-links
cd ./teapods/teapod-links && go install && cd ../../

# build frontend app
cd ./ui && pnpm install && pnpm build && cd ../

# run teatime
go run .
```

## Designdoc

- パーソナルユースなダッシュボード
- プラグインでデータスキーマの定義やアクションを記述できる
- teatime 自体はプラグインをコントロールする
- mongodb にデータを保存する
