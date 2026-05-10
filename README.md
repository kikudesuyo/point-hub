# 🚀 Point Hub

統合ポイント管理ツール (Unified Transit Point Aggregator)

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![Status](https://img.shields.io/badge/Status-Development-orange?style=for-the-badge)](https://github.com/kikudesuyo/point-hub)

PASMOを対象とした首都圏の私鉄各社のポイントを、コマンドライン一つで集約・管理するためのツールです。

## ✨ 特徴 (Features)

- **一括取得**: 複数の鉄道事業者のポイント残高と有効期限を一度に取得。
- **自動セッション管理**: 東急などのセッションを自動的に追記・更新し、再ログインの手間を軽減。
- **統一フォーマット**: 各社バラバラなデータ形式を、見やすい統一されたレポート形式で出力。

## 🚉 対応プロバイダー (Supported Providers)

| プロバイダー   | サービス名                                                  | 認証方式             | 備考                                     |
| :------------- | :---------------------------------------------------------- | :------------------- | :--------------------------------------- |
| **東京メトロ** | [メトポ (Metpo)](docs/tokyo_metro.md)                       | メール・パスワード   | [詳細](docs/tokyo_metro.md)              |
| **都営地下鉄** | [ToKoPo (トコポ)](docs/toei_metro.md)                       | 会員番号・パスワード | [詳細](docs/toei_metro.md)               |
| **東急電鉄**   | [東急ポイント](docs/tokyu.md)                               | セッショントークン   | [詳細](docs/tokyu.md)                    |
| **相模鉄道**   | [相鉄ポイント](docs/sotetsu.md)                             | メール・パスワード   | [詳細](docs/sotetsu.md)                  |
| **京浜急行**   | [京急ポイント](docs/keikyu.md)                               | 会員番号・パスワード | [詳細](docs/keikyu.md)                   |

## 🛠 セットアップ (Setup)

### 1. リポジトリのクローン

```bash
git clone https://github.com/kikudesuyo/point-hub.git
cd point-hub/api
```

### 2. 環境変数の設定

`.env.example`（もしあれば）をコピーし、各社の認証情報を入力します。

```bash
cp .env.example .env
# 各社のID/パスワードを設定
```

### 3. 依存関係のインストール

```bash
go mod tidy
```

## 🚀 使い方 (Usage)

ポイント情報を収集して表示します。

```bash
go run main.go
```

### 出力例

```text
各社ポイント収集中...

========================================
合計ポイント: 1250 pt
========================================
- Tokyo Metro (Metpo):    450 pt (最短失効: 2027-03-31)
    └ 2027-03-31: 450 pt
- Toei Metro (Tokopo):    100 pt
- Tokyu               :    500 pt
- Sotetsu             :    200 pt (最短失効: 2027-01-15)
========================================
```
