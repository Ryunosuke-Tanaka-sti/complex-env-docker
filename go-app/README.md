# Go環境

## ダミーのAPIを作成する

```bash
#立ち上げ
go mod init gin-api-project
go get -u github.com/gin-gonic/gin

#ソースコード記載

# 依存関係の整理
go mod tidy

#スタート
go run main.go

#CORS回避
go get github.com/gin-contrib/cors

```

## 拡張機能
- https://zenn.dev/tomi/articles/2020-10-22-go-docker

