# プロジェクト概要

このプロジェクトは、複数のサービスをDocker Composeで管理する複合環境を構築しています。以下は主要なサービスの概要です。

## サービス一覧

| サービス名         | 概要                                   | ポート番号                | 備考                                   |
|--------------------|----------------------------------------|--------------------------|----------------------------------------|
| next              | フロントエンドアプリケーション         | 3000                     | Next.js, TypeScript, Chakra UI         |
| go-app            | バックエンドAPI                        | 8080                     | Go, Ginフレームワーク                  |
| keycloak          | 認証・認可を管理するIDプロバイダー     | 18080 (管理コンソール)   | 管理者ユーザー: `admin`                |
| postgres          | データベースサービス                   | 5432                     | 初期化スクリプト: `init.sql`           |
| minio             | オブジェクトストレージサービス         | 9000 (API), 9001 (管理)  | 初期データ: `init_data`                |
| prefect           | ワークフロー管理ツール                 | -                        | Pythonベース                           |
| kong              | APIゲートウェイ                        | 8000 (プロキシ), 8001 (管理) | 依存サービス: `kong-database` (PostgreSQL) |
| opa               | ポリシーエンジン                       | 8181                     | -                                      |

```bash
# 事前起動
docker compose up -d kong-database

# kong用マイグレーション
docker compose run --rm kong kong migrations bootstrap

# 全体起動
docker compose up -d
```

### 1. Next.js アプリケーション (`next`)
- 概要: フロントエンドアプリケーション。
- ポート: `3000` (ローカル環境)
- 技術スタック: Next.js, TypeScript, Chakra UI。
- Dockerfile: `.devcontainer/frontend/Dockerfile`

### 2. Go アプリケーション (`go-app`)
- 概要: バックエンドAPIを提供するGoアプリケーション。
- ポート: `8080`
- 技術スタック: Go, Ginフレームワーク。
- Dockerfile: `.devcontainer/api/Dockerfile`

### 3. Keycloak (`keycloak`)
- 概要: 認証・認可を管理するIDプロバイダー。
- ポート: `18080` (管理コンソール)
- 環境変数:
  - 管理者ユーザー: `admin`
  - 管理者パスワード: `passwd`

### 4. PostgreSQL (`postgres`)
- 概要: データベースサービス。
- ポート: `5432`
- 初期化スクリプト: `docker/postgres/init/init.sql`

init.sql内のファイルを変更することで初期投入データが導入されます。現状はダミーのデータが入力されています。

### 5. pgAdmin4 (`pgadmin4`)
- 概要: PostgreSQLの管理ツール。
- ポート: `8081`
- 環境変数:
  - 管理者メール: `user@example.com`
  - 管理者パスワード: `password`

### 6. MinIO (`minio`)
- 概要: オブジェクトストレージサービス。
- ポート:
  - API: `9000`
  - 管理コンソール: `9001`
- 初期データ: `docker/minio/init_data`

### 7. Prefect (`prefect`)
- 概要: Pythonベースのワークフロー管理ツール。
- Dockerfile: `.devcontainer/batch/Dockerfile`

### 8. Kong API Gateway (`kong`)
- 概要: APIゲートウェイ。
- ポート:
  - プロキシ: `8000`
  - 管理: `8001`
- 依存サービス: `kong-database` (PostgreSQL)

### 9. Open Policy Agent (`opa`)
- 概要: ポリシーエンジン。
- ポート: `8181`

## ボリューム一覧
- `node_modules`: Next.jsの依存関係。
- `keycloak-store`: Keycloakのデータ。
- `db-store`: PostgreSQLのデータ。
- `minio-store`: MinIOのデータ。
- `kong-store`: Kongのデータ。

このプロジェクトは、フロントエンド、バックエンド、認証、データベース、ストレージ、APIゲートウェイなど、複数のサービスを統合して動作する環境を提供します。

## devcontainerについて

このプロジェクトでは、開発環境を効率的に構築するためにDevContainerを使用しています。DevContainerは、Visual Studio CodeとDockerを利用して、統一された開発環境を提供します。

### 主な構成

`.devcontainer` ディレクトリ内には、各サービスごとに必要な設定ファイルが含まれています。

#### 1. フロントエンド (`frontend`)
- **Dockerfile**: `.devcontainer/frontend/Dockerfile`
- **概要**: Next.jsをベースとしたフロントエンド開発環境を構築します。
- **特徴**:
  - Node.jsのバージョン管理。
  - 必要な依存関係のインストール。

#### 2. バックエンド (`api`)
- **Dockerfile**: `.devcontainer/api/Dockerfile`
- **概要**: Goを使用したバックエンドAPIの開発環境を提供します。
- **特徴**:
  - Goモジュールの管理。
  - Ginフレームワークのサポート。

#### 3. バッチ処理 (`batch`)
- **Dockerfile**: `.devcontainer/batch/Dockerfile`
- **概要**: Pythonを使用したバッチ処理やワークフロー管理ツール（Prefect）の開発環境を提供します。
- **特徴**:
  - Python環境のセットアップ。
  - 必要なライブラリのインストール。

### DevContainerの利点
- **統一された環境**: チーム全体で同じ開発環境を使用できるため、環境依存の問題を防ぎます。
- **迅速なセットアップ**: 必要なツールや依存関係がDockerイメージに含まれているため、新しいメンバーでもすぐに開発を開始できます。
- **拡張性**: 各サービスごとにカスタマイズされたDockerfileを使用して、特定の要件に対応可能です。

### 使用方法
1. Visual Studio Codeでプロジェクトを開きます。
2. DevContainer拡張機能をインストールします。
3. 「Reopen in Container」を選択して、コンテナ内で開発を開始します。

これにより、プロジェクト全体の開発環境が統一され、効率的な開発が可能になります。