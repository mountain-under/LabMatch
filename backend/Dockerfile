# Goの公式イメージを使用
FROM golang:1.21-alpine

# 作業ディレクトリを設定
WORKDIR /app

# go.modとgo.sumをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロードし、go.sumを更新
RUN go mod tidy

# 残りのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main ./cmd/main.go

# アプリケーションを実行
CMD ["./main"]
