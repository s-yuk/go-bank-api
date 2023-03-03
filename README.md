## テーブル設計
### user
| 列名 | デフォルト | NULL | 型 | キー | その他 |
|---|---|---|---|---|---|
| user_id  || NO | int          | PRI | auto_increment |
| username || NO | varchar(30)  |     |                |
| email    || NO | varchar(255) |     |                |
| password || NO | varchar(255) |     |                |

### account
| 列名 | デフォルト | NULL | 型 | キー | その他 |
|---|---|---|---|---|---|
| account_number |   | NO | int           | PRI ||
| balance        | 0 | NO | decimal(5, 2) |     ||
| user_id        |   | NO | int           | FK  ||

### history
| 列名 | デフォルト | NULL | 型 | キー | その他 |
|---|---|---|---|---|---|
| history_id   || NO | int           | PRI ||
| deposit      ||    | decimal(5, 2) |     ||
| transfer     ||    | decimal(5, 2) |     ||
| user_id      || NO | int           | FK  ||
| to_user_id   ||    | int           |     ||
| from_user_id ||    | int           |     ||
| created_at   || NO | timestamp     |     ||

## API設計
baseURL = /api/v1

全件取得  
GET /bank/user

一件取得  
GET /bank/user/:id

ユーザー作成  
POST /bank/user

口座取得

預金

送金

取引履歴取得