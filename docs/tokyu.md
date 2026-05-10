# 東急電鉄 (TOKYU POINT)

東急電鉄のポイントサービス「TOKYU POINT」の認証および取得方法について解説します。

## 認証方法

東急のマイページは強力なボット対策や動的なセッション管理が行われているため、通常のID/パスワードによる自動ログインではなく、ブラウザから取得したセッション情報を利用します。

### 必要な情報
- **セッショントークン**: ブラウザのCookieに含まれるトークン群
- **TOKYU_SESSION_TOKEN**: 環境変数に設定する主要なトークン

### 手順
1. [TOKYU POINT Webサービス](https://plus.tokyu.co.jp/)にブラウザでログインします。
2. デベロッパーツール等を用いて、Cookieから `__Host-plus.sessionToken` の値を取得します。
3. 取得した値を環境変数 `TOKYU_SESSION_TOKEN` に設定します。
4. 本ツール実行後、セッションが更新された場合は自動的に `.env` や `tokyu_session.json` に反映されます。

## ポイント取得の仕組み

### 技術的詳細
- **データ取得**: `https://plus.tokyu.co.jp/my/point/detail` を `GET` します。
- **解析方式**: このページは Next.js で構築されており、データは **React Server Components (RSC) Payload** としてHTML内に埋め込まれています。
- **抽出ロジック**: `self.__next_f.push` の呼び出しを正規表現で抽出し、その中から `pointBalances` を含むJSONオブジェクトをパースしてポイントと有効期限を取得します。
- **セッション維持**: Cookieのドメイン（`plus.tokyu.co.jp` と `.tokyu.co.jp`）を適切に切り分けて管理し、セッションを継続させています。
