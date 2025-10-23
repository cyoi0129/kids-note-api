# キッズノート API

子どもの TODO リストのバックエンド

# 主要アーキテクチャ

- golang
- postgre
- redis
- docker

## ローカル起動方法

### env ファイルの作成

/go/sample.env をコピーして.env ファイルを作成

### ローカルビルド

```bash
docker compose build
```

### ローカル起動

```bash
docker compose up
```

### 一時データの削除

```bash
docker compose down
```

## ビルド

```bash
go build -tags netgo -ldflags '-s -w' -o app
```
