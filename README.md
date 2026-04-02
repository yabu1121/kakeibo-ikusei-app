# kakeibo-ikusei-app

### 要件

### api 

### テーブル設計

### golang

domain => repository => persistense => interface => usecase 

├── cmd/
│   └── server/
│       └── main.go           # 全ての層を繋ぎ合わせる（DI）場所
├── domain/                   # 【最内部】ビジネスルール・魂
│   ├── model/                # 構造体（User, Expense, Character）
│   │   ├── expense.go        # 支出の定義
│   │   └── character.go      # キャラの定義 + CalcExpメソッド
│   └── repository/           # インターフェース（抽象）
│       ├── expense_repo.go   # 「保存する」という約束事だけ定義
│       └── character_repo.go
├── usecase/                  # 【中間】アプリケーションの台本
│   ├── expense_usecase.go    # 「支出を保存してキャラを更新する」手順
│   └── character_usecase.go
├── interfaces/               # 【外部接続】Web（Echo）の入り口
│   ├── handler/              # リクエストを受け取り、Usecaseを呼ぶ
│   │   ├── expense_handler.go
│   │   └── character_handler.go
│   └── response/             # フロントに返す用の型（DTO）
└── infrastructure/           # 【詳細】SQLite、GORMの実装
    └── persistence/          # 具体的なDB操作（SQL発行）
        └── sqlite/
            ├── db.go         # SQLite接続設定
            ├── expense.go    # GORMを使った具体的な保存処理
            └── character.go
