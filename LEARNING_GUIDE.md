# OSS リリース学習ガイド

SPEC.mdに基づいて、実際にOSSをリリースする方法を学習するための詳細な手順書です。

## 📋 前提条件

以下のツールがインストールされていることを確認してください：

- Node.js (v14以上)
- npm
- Docker
- Go (1.20以上)
- Git
- GitHub CLI (`gh` コマンド)

```bash
# GitHub CLIのインストール（macOS）
brew install gh

# GitHub CLIでログイン
gh auth login
```

## 1️⃣ NPMパッケージを公開する方法

### 1.1 NPMアカウントの準備

1. [npmjs.com](https://www.npmjs.com/) でアカウントを作成
2. メールアドレスを認証
3. ターミナルでnpmにログイン：

```bash
npm login
# Username, Password, Email, OTP（2FA有効時）を入力
```

### 1.2 パッケージのテスト

```bash
cd npm-package

# パッケージの動作確認
node test.js

# パッケージ情報の確認
npm pack --dry-run
```

### 1.3 テスト公開（推奨）

本番公開前にテストすることを推奨：

```bash
# package.jsonのnameを一時的に変更
# 例: "@your-username/test-string-formatter"

# テスト公開
npm publish --dry-run  # ドライラン
npm publish --access public  # 実際に公開
```

### 1.4 公開の確認

```bash
# 公開されたパッケージを確認
npm view @your-username/test-string-formatter

# 別ディレクトリでインストールテスト
mkdir /tmp/test-npm && cd /tmp/test-npm
npm init -y
npm install @your-username/test-string-formatter
```

### 1.5 パッケージの非公開化（unpublish）

**重要**: 学習後は必ず非公開にしましょう

```bash
# 72時間以内なら完全削除可能
npm unpublish @your-username/test-string-formatter --force

# または特定バージョンのみ削除
npm unpublish @your-username/test-string-formatter@1.0.0
```

**注意事項**:
- 一度公開したパッケージ名は、削除後24時間は再利用できません
- 他のパッケージが依存している場合は削除できません
- npmのポリシーにより、公開から72時間を超えると削除が難しくなります

## 2️⃣ DockerイメージをGitHub Container Registryで公開する方法

### 2.1 GitHub Personal Access Token (PAT) の作成

1. GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
2. "Generate new token (classic)" をクリック
3. 以下のスコープを選択：
   - `write:packages` - パッケージのアップロード
   - `read:packages` - パッケージの読み取り
   - `delete:packages` - パッケージの削除（クリーンアップ用）
4. トークンを安全に保存

### 2.2 GitHub Container Registryへのログイン

```bash
# 環境変数にトークンを設定
export GITHUB_TOKEN=your_token_here

# ghcr.ioにログイン
echo $GITHUB_TOKEN | docker login ghcr.io -u YOUR_GITHUB_USERNAME --password-stdin
```

### 2.3 Dockerイメージのビルドとタグ付け

```bash
cd docker-app

# イメージをビルド
docker build -t oss-demo-app .

# ghcr.io用にタグ付け
docker tag oss-demo-app ghcr.io/YOUR_GITHUB_USERNAME/oss-demo-app:latest
docker tag oss-demo-app ghcr.io/YOUR_GITHUB_USERNAME/oss-demo-app:1.0.0

# ローカルでテスト実行
docker run -p 3000:3000 ghcr.io/YOUR_GITHUB_USERNAME/oss-demo-app:latest
# ブラウザで http://localhost:3000 を確認
```

### 2.4 イメージのプッシュ

```bash
# GitHub Container Registryにプッシュ
docker push ghcr.io/YOUR_GITHUB_USERNAME/oss-demo-app:latest
docker push ghcr.io/YOUR_GITHUB_USERNAME/oss-demo-app:1.0.0
```

### 2.5 イメージの公開設定

1. GitHubプロフィール → Packages タブ
2. `oss-demo-app` パッケージをクリック
3. Package settings → Change visibility → Public

### 2.6 公開イメージのテスト

```bash
# 別のターミナルやマシンで（ログインなしで）
docker pull ghcr.io/YOUR_GITHUB_USERNAME/oss-demo-app:latest
docker run -p 3000:3000 ghcr.io/YOUR_GITHUB_USERNAME/oss-demo-app:latest
```

### 2.7 イメージの削除（学習後）

```bash
# 特定のタグを削除
docker rmi ghcr.io/YOUR_GITHUB_USERNAME/oss-demo-app:1.0.0
docker rmi ghcr.io/YOUR_GITHUB_USERNAME/oss-demo-app:latest

# GitHub上から完全削除
# GitHubのPackagesページから削除、または：
gh api -X DELETE /user/packages/container/oss-demo-app/versions/VERSION_ID
```

## 3️⃣ ビルドしたバイナリをGitHub Releasesページで公開する方法

### 3.1 Goバイナリのビルド

```bash
cd go-cli

# 現在のプラットフォーム用にビルド
go build -o oss-demo-cli

# クロスプラットフォームビルド
make build-all

# ビルド結果を確認
ls -la dist/
```

### 3.2 GitHubリポジトリの準備

```bash
# リポジトリをGitHubに作成（まだの場合）
gh repo create study-oss-release-test --public --clone

# 既存のコードをコミット
git add .
git commit -m "Add release demo projects"
git push origin main
```

### 3.3 GitHub Releasesの作成（CLI版）

```bash
# タグを作成
git tag v1.0.0
git push origin v1.0.0

# リリースを作成
gh release create v1.0.0 \
  --title "Release v1.0.0" \
  --notes "初回リリース
  
  ## 含まれるもの
  - macOS (Intel/Apple Silicon)
  - Linux (AMD64/ARM64)  
  - Windows (AMD64)
  
  ## インストール方法
  お使いのOSに合わせたバイナリをダウンロードしてください。" \
  --draft

# バイナリをアップロード
cd go-cli/dist
gh release upload v1.0.0 *.* --clobber

# ドラフトを公開
gh release edit v1.0.0 --draft=false
```

### 3.4 GitHub Releasesの作成（Web UI版）

1. GitHubリポジトリページを開く
2. 右側の "Releases" をクリック
3. "Draft a new release" をクリック
4. 設定項目：
   - **Choose a tag**: `v1.0.0` を入力して "Create new tag"
   - **Release title**: `Release v1.0.0`
   - **Describe this release**: リリースノートを記載
   - **Attach binaries**: ビルドしたファイルをドラッグ&ドロップ
5. "Save draft" で下書き保存、または "Publish release" で公開

### 3.5 リリースの確認とダウンロードテスト

```bash
# リリース一覧を確認
gh release list

# リリースの詳細を確認
gh release view v1.0.0

# バイナリをダウンロード（別ディレクトリで）
mkdir /tmp/test-release && cd /tmp/test-release
gh release download v1.0.0 --repo YOUR_GITHUB_USERNAME/study-oss-release-test

# または curl でダウンロード
curl -LO https://github.com/YOUR_GITHUB_USERNAME/study-oss-release-test/releases/download/v1.0.0/oss-release-demo-cli-darwin-arm64
chmod +x oss-release-demo-cli-darwin-arm64
./oss-release-demo-cli-darwin-arm64 --version
```

### 3.6 リリースの削除（学習後）

```bash
# リリースを削除（タグは残る）
gh release delete v1.0.0 --yes

# タグも削除
git push --delete origin v1.0.0
git tag -d v1.0.0
```

## 4️⃣ リリース作業の自動化

### 4.1 GitHub Actionsの動作確認

```bash
# ワークフローの一覧確認
gh workflow list

# 手動実行（workflow_dispatch）
gh workflow run release-all.yml \
  -f release_version=1.0.0 \
  -f npm_publish=false \
  -f docker_publish=true \
  -f go_release=true

# 実行状況の確認
gh run list
gh run watch
```

### 4.2 自動リリースのテスト

```bash
# リリースタグでトリガー
git tag release-1.0.1
git push origin release-1.0.1

# Actions の実行を確認
gh run list --workflow=release-all.yml
```

## 🧹 学習後のクリーンアップ

### 完全クリーンアップチェックリスト

```bash
# 1. NPMパッケージの削除
npm unpublish @your-username/test-package --force

# 2. Docker イメージの削除
# GitHub Packages ページから削除

# 3. GitHub Releases の削除
gh release list | awk '{print $1}' | xargs -I {} gh release delete {} --yes

# 4. すべてのタグを削除
git tag | xargs -I {} git push --delete origin {}
git tag | xargs git tag -d

# 5. テストリポジトリの削除（必要に応じて）
gh repo delete YOUR_GITHUB_USERNAME/study-oss-release-test --yes
```

## ⚠️ 注意事項

1. **NPMの注意点**
   - テスト用パッケージは必ず削除する
   - スコープ（@username/）を付けて名前の衝突を避ける
   - 72時間以内に削除することを推奨

2. **Docker Registryの注意点**
   - Personal Access Tokenは安全に管理
   - 不要なイメージは定期的に削除
   - タグは意味のある名前を使用

3. **GitHub Releasesの注意点**
   - プレリリース機能を活用してテスト
   - チェックサムファイルを含めてセキュリティ向上
   - リリースノートは詳細に記載

## 📚 追加学習リソース

- [npm Documentation](https://docs.npmjs.com/)
- [GitHub Packages Documentation](https://docs.github.com/en/packages)
- [GitHub Releases Documentation](https://docs.github.com/en/repositories/releasing-projects-on-github)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

## 🎯 学習のゴール

この手順を完了すると、以下ができるようになります：

✅ NPMパッケージの公開と管理
✅ Dockerイメージの配布
✅ バイナリファイルの配布
✅ CI/CDパイプラインの構築
✅ 自動リリースプロセスの実装