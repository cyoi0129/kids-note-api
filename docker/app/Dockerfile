FROM golang:latest

# コンテナ作業ディレクトリの変更
WORKDIR /go/src/app/go
# ホストOSの ./go の中身を作業ディレクトリに追加
ADD ./go .

# パッケージのインストール
RUN go install github.com/air-verse/air@latest
RUN go mod download

CMD ["air", "-c", ".air.toml"]