# 都営地下鉄 (ToKoPo / トコポ)

都営交通のポイントサービス「ToKoPo」の認証および取得方法について解説します。

## 認証方法

### 必要な情報
- **会員番号 (Card No)**: 会員カードに記載された番号
- **パスワード (Password)**: 登録時に設定したパスワード

### 手順
1. [ToKoPo公式サイト](https://www.kotsu.metro.tokyo.jp/tokopo/)で会員登録を行います。
2. 登録後、駅に設置されているポイントチャージ機でICカードの登録を行う必要があります。
3. 会員番号とパスワードを環境変数 `TOEI_USER_ID` と `TOEI_METRO_PASSWORD` に設定します。

## ポイント取得の仕組み

### 技術的詳細
- **認証**: `https://tokopo.jp/gv/pc/login/PcLoginLoginAction.do` に対して `POST` リクエストを送信し、セッションCookieを取得します。
- **データ取得**: マイページ (`MyMenuInitAction.do`) を `GET` し、HTMLをパースします。
- **取得項目**:
    - 保有ポイント
- **エンコーディング**: サイトが `Shift-JIS` を使用しているため、取得時に `UTF-8` に変換して処理しています。
