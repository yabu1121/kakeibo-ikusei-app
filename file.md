.
├── ./CLEAN_ARCHITECTURE.md
├── ./README.md
├── ./backend
│   ├── ./backend/cmd
│   │   └── ./backend/cmd/server
│   │       ├── ./backend/cmd/server/game.db
│   │       └── ./backend/cmd/server/main.go
│   ├── ./backend/domain
│   │   ├── ./backend/domain/model
│   │   │   ├── ./backend/domain/model/category.go
│   │   │   ├── ./backend/domain/model/character.go
│   │   │   ├── ./backend/domain/model/expense.go
│   │   │   ├── ./backend/domain/model/timestamps.go
│   │   │   └── ./backend/domain/model/user.go
│   │   ├── ./backend/domain/repository
│   │   │   ├── ./backend/domain/repository/category.go
│   │   │   ├── ./backend/domain/repository/character.go
│   │   │   ├── ./backend/domain/repository/expense.go
│   │   │   └── ./backend/domain/repository/notifier.go
│   │   └── ./backend/domain/service
│   │       └── ./backend/domain/service/exp.go
│   ├── ./backend/go.mod
│   ├── ./backend/go.sum
│   ├── ./backend/handler
│   │   ├── ./backend/handler/category.go
│   │   ├── ./backend/handler/character.go
│   │   ├── ./backend/handler/expense.go
│   │   └── ./backend/handler/slack.go
│   ├── ./backend/infrastructure
│   │   ├── ./backend/infrastructure/db.go
│   │   ├── ./backend/infrastructure/persistence
│   │   │   └── ./backend/infrastructure/persistence/sqlite
│   │   │       ├── ./backend/infrastructure/persistence/sqlite/category.go
│   │   │       ├── ./backend/infrastructure/persistence/sqlite/character.go
│   │   │       └── ./backend/infrastructure/persistence/sqlite/expense.go
│   │   └── ./backend/infrastructure/slack
│   │       └── ./backend/infrastructure/slack/notifier.go
│   └── ./backend/usecase
│       ├── ./backend/usecase/category.go
│       ├── ./backend/usecase/character.go
│       ├── ./backend/usecase/expense.go
│       └── ./backend/usecase/slack_notifier.go
├── ./file.md
└── ./frontend
    ├── ./frontend/AGENTS.md
    ├── ./frontend/CLAUDE.md
    ├── ./frontend/README.md
    ├── ./frontend/app
    │   ├── ./frontend/app/favicon.ico
    │   ├── ./frontend/app/globals.css
    │   ├── ./frontend/app/layout.tsx
    │   └── ./frontend/app/page.tsx
    ├── ./frontend/eslint.config.mjs
    ├── ./frontend/next-env.d.ts
    ├── ./frontend/next.config.ts
    ├── ./frontend/package-lock.json
    ├── ./frontend/package.json
    ├── ./frontend/postcss.config.mjs
    ├── ./frontend/public
    │   ├── ./frontend/public/file.svg
    │   ├── ./frontend/public/globe.svg
    │   ├── ./frontend/public/next.svg
    │   ├── ./frontend/public/vercel.svg
    │   └── ./frontend/public/window.svg
    └── ./frontend/tsconfig.json

16 directories, 49 files
