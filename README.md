# TASCHOLA

TSSCHOLA is a simple todo list manager which can sync with Tokyo Tech T2SCHOLA.

## Set Up

1. Clone this repository
2. `docker-compose up -d`

## For Developers

- Docker Compose について

### Frontend

本プロジェクトにて採用した記述等に関する説明

- React + TypeScript

  Next.js を使用しなかった理由

### Backend

- CORS について

- TablePlus について

- MySQL との接続について

  `wait-for-it.sh`を用いて、MySQL が起動するまで待機するようにしていたが、`healthcheck:`を利用することにより、`docker-compose.yml`だけで依存関係と実行順番制御が行える。

  - heath check について (死活監視)

    Go Gin においては、`/health`エンドポイントを作成し、問題なく Gin が起動している場合は`200`を返すようにしている。

    また、MySQL については mysqladmin ping で問題なく接続できるかいなかを確認している。

  - health check の詳細

    `docker-compose.yml`の `healthcheck:`, `test:` に実際に確かめるために使用しているコマンドが記載されている。

    また、health check の周期なども`docker-compose.yml`で設定している。

### DB

### 参考資料

#### Frontend

#### Backend

- [Go Gin CORS + 認証](https://qiita.com/bty__/items/f8c4393bd7701a1d703c)
- [docker-compose におけるヘルスチェック](https://qiita.com/hichika/items/9b96634d471246359e66)

#### DB
