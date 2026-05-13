# kakeibo-ikusei-app

### 制作背景  
従来の家計簿アプリは作業化している。これに関しては問題はないが、出費を記録するのを忘れてしまった場合、急激に家計簿をつける意義を見失ってしまうなと考えていました。また、自分が使っている家計簿のアプリは出費を登録するのにかなり手順が多すぎて面倒くさいということがありました。 このアプリではそれら二つの課題を解決したいということで自分用の家計簿アプリを作成しようと思っています。  
フロントエンドの部分を効率化して簡単に入力できるようにする。  
大きな出費ほど記録を忘れないようなロジックを持たせる。  
小さな出費でも家計簿のことを一日一回程度は思い出させて記入させる。  
というようなことを考えるとキャラクター育成などを用いると要件を満たせると思ったので
育成　×　出費記録のkakebonを作ることにしました。

### 技術選定
client: Next.js
server: Golang

言語選定理由:
Next.js: サーバーサイドでも、クライアントサイドでも実装できるため、速度向上、SEO向上に貢献できると考えたため。
Golang: Cppに次いで高速な言語なため、UXに対して貢献できるため、apiサーバーとして用いる際に並列処理を得意としているため、パフォーマンス向上が見込めるため。


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

### 実装済みapi
getAllCategory (すべてのカテゴリー取得)

CreateCategory (カテゴリー作成)

GetCharacterInfomation (キャラクターの情報取得)

LoginBonus (ログインボーナスのキャラクター経験値付与関数)

RecordExpense （消費を記録する関数)

GetAllExpense (消費を人関係なしに取得する関数)

GetByID (消費をIDで取得する関数)

DeleteByID (消費をIDで削除、ポイント付与取り上げ関数)

Notify (メッセージを受け取ってslackに送信する関数)

SignUp (キャラクター)

GetByEmail (emailでユーザが存在するかを返す関数)

Login (jwt発行)

UpdateExpense (消費を経験値がらみでアップデートする)
