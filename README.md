# OSS Release Study Repository

このリポジトリは、GitHub上でOSSをリリースする方法を学習するための実践的なサンプル集です。

## 📚 学習内容

### 1. NPMパッケージの公開
- **場所**: `/npm-package`
- **内容**: シンプルな文字列フォーマッタライブラリ
- **学習ポイント**:
  - package.jsonの設定
  - npmへの公開手順
  - スコープパッケージの管理

### 2. Dockerイメージの公開
- **場所**: `/docker-app`
- **内容**: Node.jsベースのWebアプリケーション
- **学習ポイント**:
  - Dockerfileの作成
  - GitHub Container Registry (ghcr.io)への公開
  - マルチステージビルド

### 3. バイナリ配布
- **場所**: `/go-cli`
- **内容**: Go言語で実装されたCLIツール
- **学習ポイント**:
  - クロスプラットフォームビルド
  - GitHub Releasesでの配布
  - チェックサムの生成

### 4. リリース自動化
- **場所**: `/.github/workflows`
- **内容**: GitHub Actionsワークフロー
- **学習ポイント**:
  - CI/CDパイプライン構築
  - 自動リリースプロセス
  - 複数コンポーネントの統合リリース

## 🚀 使い方

### NPMパッケージのテスト
```bash
cd npm-package
npm test
```

### Dockerイメージのローカルビルド
```bash
cd docker-app
docker build -t demo-app .
docker run -p 3000:3000 demo-app
```

### Go CLIのビルド
```bash
cd go-cli
make build-all
```

## 📦 リリース方法

### 手動リリース

#### NPMパッケージ
```bash
cd npm-package
npm login
npm publish --access public
```

#### Dockerイメージ
```bash
cd docker-app
docker build -t ghcr.io/USERNAME/app:latest .
docker push ghcr.io/USERNAME/app:latest
```

#### Goバイナリ
```bash
cd go-cli
make build-all
# GitHub Releasesページで手動アップロード
```

### 自動リリース

タグをプッシュすることで自動リリースが実行されます：

```bash
git tag release-1.0.0
git push origin release-1.0.0
```

または、GitHub ActionsのUIから手動でワークフローを実行できます。

## 📝 必要な設定

### Secrets設定
- `NPM_TOKEN`: NPMの認証トークン
- `GITHUB_TOKEN`: 自動で提供される（設定不要）

### Permissions設定
- Packages: write権限
- Contents: write権限（Releasesの作成用）

## 🏗️ プロジェクト構造

```
.
├── npm-package/          # NPMパッケージ
│   ├── index.js
│   ├── package.json
│   └── README.md
├── docker-app/          # Dockerアプリケーション
│   ├── app.js
│   ├── Dockerfile
│   └── package.json
├── go-cli/              # Go CLIツール
│   ├── main.go
│   ├── go.mod
│   └── Makefile
└── .github/workflows/   # GitHub Actions
    ├── npm-publish.yml
    ├── docker-publish.yml
    ├── go-release.yml
    └── release-all.yml
```

## 📄 ライセンス

MIT License

## 🤝 コントリビューション

学習用リポジトリのため、プルリクエストは受け付けていません。
フォークして自由に実験してください。