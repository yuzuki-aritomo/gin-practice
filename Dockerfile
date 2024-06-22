# Go の公式イメージを使用
FROM golang:1.22.4

# 作業ディレクトリを作成
WORKDIR /app

# Go モジュールファイルをコピー
COPY go.mod go.sum ./

# 依存関係をインストール
RUN go mod download

# プロジェクトファイルをコピー
COPY . .

# Air をインストール
RUN go install github.com/air-verse/air@latest

# Air を起動
ENTRYPOINT ["air"]
