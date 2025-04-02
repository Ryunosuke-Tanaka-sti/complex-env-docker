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

#CORS回避:Gin用のCors設定
go get github.com/gin-contrib/cors
go get github.com/gin-contrib/cors@none #パッケージ削除の方法

#DB関連パッケージ
go get gorm.io/driver/postgres
go get gorm.io/gorm

```

### Corsの設定
```go
router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000", "https://example.com"}, // 許可するオリジン
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                 // 許可するHTTPメソッド
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},      // 許可するヘッダー
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,           // Cookie などの認証情報を許可
    MaxAge:           12 * time.Hour, // プリフライトリクエストのキャッシュ時間
}))
```

### ディレクトリ構造
```bash
├── main.go                # アプリケーションのエントリーポイント
├── config                 # 設定関連
│   └── database.go        # データベース接続設定
├── controllers            # コントローラー（ビジネスロジック）
│   └── user_controller.go # usersテーブル用のCRUD操作ハンドラ
├── models                 # データモデル
│   └── user.go            # usersテーブルのモデル定義
├── routes                 # ルーティング設定
│   └── routes.go          # APIエンドポイントのルート定義
├── database               # マイグレーションやシードデータ関連
│   └── migration.go       # テーブル作成・更新処理
└── go.mod                 # Goモジュールファイル
```
Gormでは、modelsで定義したファイルの複数形のテーブルが参照される？

対象とするDBは立ち上げ時点でマイグレーションされているdummy_db

参照 ../docker/postgres/init/init.sql
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```


## 拡張機能
- https://zenn.dev/tomi/articles/2020-10-22-go-docker

