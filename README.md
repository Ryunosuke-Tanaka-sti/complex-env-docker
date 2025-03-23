# 構築

- keycloak
  - [https://quay.io/repository/keycloak/keycloak?tab=tags]
  - [https://zenn.dev/kazuma_r5/articles/d464cd0bf380da]
- postgress
  - [https://hub.docker.com/_/postgres/tags]
  - [https://zenn.dev/re24_1986/articles/b76c3fd8f76aec]
  - PORT:5432
- pgadmin4
  - [https://hub.docker.com/r/dpage/pgadmin4/]
  - [楽しようぜ](https://zenn.dev/onozaty/articles/postgresql-pgadmin-container)
  - PORT:8080
- cauchbase
  - 二つのイメージの差異は何？：違いは特になし
    - [https://hub.docker.com/_/couchbase]
    - [https://hub.docker.com/r/couchbase/server]
  - [https://dzone.com/articles/couchbase-cluster-using-docker-compose]
  - メモ
    - bindマウントでも非rootユーザーで権限で保存される親切設計
  - PORT:8091・11210
- MinIO
  - [MinIO Docker Image](https://hub.docker.com/r/minio/minio)
  - [https://github.com/minio/minio/blob/master/docs/docker/README.md]
  - PORT:9000・9001
- MinIO mc
  - Client Toolのセット構築
  - [MinIO mc Docker Image](https://hub.docker.com/r/minio/mc)
  - Client Toolとの切り分け

    | 特徴     | MinIO (minio)                  | MinIO Client (minio/mc)                  |
    | -------- | ------------------------------ | ---------------------------------------- |
    | 役割     | オブジェクトストレージサーバー | クライアントツール                       |
    | 機能     | データ保存・管理               | ストレージ操作（作成、コピー、削除など） |
    | 用途     | S3互換ストレージの構築・運用   | サーバーやバケットへの操作               |
    | 起動方法 | サーバープロセスとして起動     | CLIツールとして利用                      |
