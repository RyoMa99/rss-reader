# RSS Reader

`src/constant/blog_list.go`ファイル内で管理しているURLからRSSを取得し、実行時から`config.dev.yaml`の`range`時間分前の間に追加されたコンテンツのタイトルとURLをまとめる。

## コマンド実行
```bash
cd src
export $(cat dev.env)
go run main.go
```

## docker image build
```bash
docker build -t rss-reader .
```

## docker image run
```bash
docker run --mount type=bind,source="$(pwd)"/config.dev.yaml,target=/config/config.yaml rss-reader
```
