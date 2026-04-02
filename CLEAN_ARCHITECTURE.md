# クリーンアーキテクチャ解説

## なぜクリーンアーキテクチャ？

コードが成長するにつれ「ハンドラーにDB処理が直書きされている」「テストを書こうとしたらDBが必要になった」という問題が起きる。クリーンアーキテクチャはこれを防ぐための**依存関係のルール**。

---

## 全体像：依存の向き

```
┌──────────────────────────────────┐
│         Handler (外側)           │  ← HTTPリクエスト/レスポンス
│  ┌────────────────────────────┐  │
│  │       Usecase              │  │  ← アプリのビジネスロジック
│  │  ┌──────────────────────┐  │  │
│  │  │  Domain (最内側)     │  │  │  ← エンティティ・インターフェース
│  │  └──────────────────────┘  │  │
│  └────────────────────────────┘  │
│                                  │
│  Infrastructure (外側)           │  ← DBの実装
└──────────────────────────────────┘
```

**大原則: 矢印は必ず内側に向かう。内側は外側を知らない。**

---

## このプロジェクトのディレクトリ構成

```
backend/
├── cmd/server/
│   └── main.go                  # DIの配線のみ。各層を組み立てる起点
│
├── domain/                      # 最内側。外部への依存ゼロ
│   ├── model/                   # エンティティ（純粋なデータ構造）
│   │   ├── expense.go           # Expense構造体
│   │   ├── character.go         # Character構造体
│   │   ├── category.go
│   │   ├── user.go
│   │   └── timestamps.go
│   ├── repository/              # リポジトリの「契約」（interfaceのみ）
│   │   ├── expense.go           # ExpenseRepository interface
│   │   └── character.go         # CharacterRepository interface
│   └── service/                 # ドメインロジック（純粋な計算）
│       └── exp.go               # CalcExp: DB・HTTPを一切使わない
│
├── usecase/                     # アプリケーション層
│   ├── expense.go               # 支出記録・一覧取得のユースケース
│   └── character.go             # キャラクター取得・ログインボーナスのユースケース
│
├── handler/                     # プレゼンテーション層
│   ├── expense.go               # ExpenseHandler + ExpenseRequest
│   └── character.go             # CharacterHandler + CharacterResponse
│
└── infrastructure/              # 最外側。具体的な技術の実装
    ├── db.go                    # DB接続の初期化
    └── persistence/
        └── sqlite/
            ├── expense.go       # ExpenseRepository の実装
            └── character.go     # CharacterRepository の実装
```

---

## 各層の責務

### domain/model — エンティティ

純粋なデータ構造。GORMのタグは持つが、HTTPやDBの処理ロジックは持たない。

```go
// OK: データ構造の定義
type Character struct {
    ID             string
    CurrentLevel   int
    CurrentExp     int
    ExpToNextLevel int
}

// NG: ここにHTTPリクエストの型を置かない
// type CharacterResponse struct { ... }  ← handler層の責務
```

### domain/repository — インターフェース（契約）

「こういう操作ができること」という約束だけを定義。実装は知らない。

```go
type CharacterRepository interface {
    GetByUserId(userID string) (*model.Character, error)
    Update(character *model.Character) error
}
```

usecase層はこのinterfaceだけに依存するため、SQLiteをPostgreSQLに変えても usecase のコードは一切変わらない。

### domain/service — ドメインサービス

複数のエンティティにまたがる、または単体のエンティティに収まらない純粋なビジネスロジック。

```go
// DBもHTTPも一切使わない純粋な計算
func CalcExp(char *model.Character, exp int) {
    char.CurrentExp += exp
    for char.CurrentExp >= char.ExpToNextLevel {
        // レベルアップ処理...
    }
}
```

### usecase — アプリケーション層

「支出を記録する」「ログインボーナスを付与する」といった1つのユースケースを実現する。repositoryとdomain serviceを組み合わせる。

```go
func (u *ExpenseUsecase) RecordExpense(expense *model.Expense) (*model.Character, error) {
    // 1. 支出をDBに保存（repositoryに委譲）
    u.expenseRepo.Create(expense)
    // 2. キャラクターを取得（repositoryに委譲）
    char, _ := u.characterRepo.GetByUserId(expense.UserID)
    // 3. 経験値計算（domain serviceに委譲）
    service.CalcExp(char, expense.Amount/100)
    // 4. キャラクターを更新（repositoryに委譲）
    u.characterRepo.Update(char)
    return char, nil
}
```

### handler — プレゼンテーション層

HTTPの世界とアプリの世界をつなぐ。リクエストのパース・バリデーション・レスポンスの整形だけを担う。

```go
// リクエスト/レスポンス型はここで定義
type ExpenseRequest struct {
    Amount     int    `json:"amount"`
    Name       string `json:"name"`
}

func (h *ExpenseHandler) RecordExpense(c echo.Context) error {
    // 1. リクエストをパース
    c.Bind(&req)
    // 2. ユースケースを呼ぶ（HTTPを知らないusecaseに委譲）
    char, err := h.usecase.RecordExpense(expense)
    // 3. レスポンスを返す
    return c.JSON(http.StatusOK, char)
}
```

### infrastructure — インフラ層

domain/repositoryのinterfaceを実装する。GORMやSQLiteなど具体的な技術がここに集まる。

```go
// ExpenseRepository interfaceを実装
func (r *expenseRepository) Create(expense *model.Expense) error {
    return r.db.Create(expense).Error
}
```

### cmd/server/main.go — DIの配線

各層のインスタンスを生成して組み立てるだけ。ビジネスロジックは一切書かない。

```go
func main() {
    db := infrastructure.InitDB()

    // Infrastructure → Usecase → Handler の順に組み立てる
    expenseRepo    := sqlite.NewExpenseRepository(db)
    expenseUsecase := usecase.NewExpenseUsecase(expenseRepo, characterRepo)
    expenseHandler := handler.NewExpenseHandler(expenseUsecase)

    e.POST("/expense", expenseHandler.RecordExpense)
    e.Start(":8080")
}
```

---

## よくある間違い（このプロジェクトで起きていたこと）

| 間違い | なぜNG | 正しい場所 |
|--------|--------|-----------|
| `model/expense.go` に `ExpenseRequest` を定義 | modelはHTTPを知らない | `handler/expense.go` |
| `model/character.go` に `CharacterResponse` を定義 | 同上 | `handler/character.go` |
| `domain/service/exp.go` で `UserCharacter` を使う | パッケージ外の型を参照できない | `model.Character` を import して使う |
| `infrastructure/db.go` で旧MVPの構造体を直接使う | domain modelを使うべき | `model.User{}` 等を import |
| `main.go` にハンドラー・モデル・ロジックを全部書く | 全層が混在している | 各責務を適切な層に分割 |

---

## 依存関係の確認方法

「この import は正しいか？」を確認するには矢印の方向を見る。

```
handler  →  usecase  →  domain
                ↑
        infrastructure
```

- `handler` が `usecase` を import する → OK（外→内）
- `usecase` が `domain/repository` を import する → OK（外→内）
- `infrastructure` が `domain/model` を import する → OK（外→内）
- `domain/model` が `handler` を import する → **NG**（内が外を知ってしまう）
- `usecase` が `infrastructure` を import する → **NG**（usecaseがDBの実装に依存してしまう）

---

## 今後の拡張ポイント

### トランザクション
現状の実装では「支出の保存」と「キャラクター更新」が別トランザクション。障害時に不整合が起きる可能性がある。対策として `TransactionManager` interfaceを domain 層に定義し、infrastructure層で実装する。

### JWT認証
`handler/*.go` の `userID := "dummy-user-id"` を、JWT middlewareから取得するよう変更する。usecase・domain層は変更不要。

### テスト
usecase のテストは `repository.CharacterRepository` のモックを渡すだけで書ける。DBなしでビジネスロジックをテストできるのがクリーンアーキテクチャの恩恵。

```go
func TestLoginBonus(t *testing.T) {
    mockRepo := &MockCharacterRepository{ /* ... */ }
    uc := usecase.NewCharacterUsecase(mockRepo)
    char, err := uc.LoginBonus("user-123")
    // DBなしでテストできる
}
```
