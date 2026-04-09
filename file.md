.
├── ./CLEAN_ARCHITECTURE.md
├── ./README.md
├── ./api.http
├── ./backend
│   ├── ./backend/cmd
│   │   └── ./backend/cmd/server
│   │       ├── ./backend/cmd/server/game.db
│   │       └── ./backend/cmd/server/main.go
│   ├── ./backend/domain
│   │   ├── ./backend/domain/model
│   │   │   ├── ./backend/domain/model/assets.go
│   │   │   ├── ./backend/domain/model/category.go
│   │   │   ├── ./backend/domain/model/character.go
│   │   │   ├── ./backend/domain/model/expense.go
│   │   │   ├── ./backend/domain/model/timestamps.go
│   │   │   └── ./backend/domain/model/user.go
│   │   ├── ./backend/domain/repository
│   │   │   ├── ./backend/domain/repository/category.go
│   │   │   ├── ./backend/domain/repository/character.go
│   │   │   ├── ./backend/domain/repository/expense.go
│   │   │   ├── ./backend/domain/repository/notifier.go
│   │   │   └── ./backend/domain/repository/user.go
│   │   └── ./backend/domain/service
│   │       └── ./backend/domain/service/exp.go
│   ├── ./backend/go.mod
│   ├── ./backend/go.sum
│   ├── ./backend/handler
│   │   ├── ./backend/handler/category.go
│   │   ├── ./backend/handler/character.go
│   │   ├── ./backend/handler/expense.go
│   │   ├── ./backend/handler/slack.go
│   │   ├── ./backend/handler/user.go
│   │   └── ./backend/handler/utils
│   │       ├── ./backend/handler/utils/auth.go
│   │       └── ./backend/handler/utils/middleware.go
│   ├── ./backend/infrastructure
│   │   ├── ./backend/infrastructure/db.go
│   │   ├── ./backend/infrastructure/persistence
│   │   │   └── ./backend/infrastructure/persistence/sqlite
│   │   │       ├── ./backend/infrastructure/persistence/sqlite/category.go
│   │   │       ├── ./backend/infrastructure/persistence/sqlite/character.go
│   │   │       ├── ./backend/infrastructure/persistence/sqlite/expense.go
│   │   │       └── ./backend/infrastructure/persistence/sqlite/user.go
│   │   └── ./backend/infrastructure/slack
│   │       └── ./backend/infrastructure/slack/notifier.go
│   └── ./backend/usecase
│       ├── ./backend/usecase/category.go
│       ├── ./backend/usecase/character.go
│       ├── ./backend/usecase/expense.go
│       ├── ./backend/usecase/slack_notifier.go
│       └── ./backend/usecase/user.go
├── ./file.md
└── ./frontend
    ├── ./frontend/README.md
    ├── ./frontend/app
    │   ├── ./frontend/app/(top)
    │   │   ├── ./frontend/app/(top)/layout.tsx
    │   │   └── ./frontend/app/(top)/page.tsx
    │   ├── ./frontend/app/actions
    │   │   ├── ./frontend/app/actions/auth.ts
    │   │   ├── ./frontend/app/actions/category.ts
    │   │   ├── ./frontend/app/actions/character.ts
    │   │   ├── ./frontend/app/actions/expense.ts
    │   │   ├── ./frontend/app/actions/slack.ts
    │   │   └── ./frontend/app/actions/util.ts
    │   ├── ./frontend/app/favicon.ico
    │   ├── ./frontend/app/globals.css
    │   ├── ./frontend/app/layout.tsx
    │   ├── ./frontend/app/login
    │   │   └── ./frontend/app/login/page.tsx
    │   ├── ./frontend/app/record
    │   │   └── ./frontend/app/record/page.tsx
    │   └── ./frontend/app/signup
    │       └── ./frontend/app/signup/page.tsx
    ├── ./frontend/components
    │   ├── ./frontend/components/features
    │   │   ├── ./frontend/components/features/LoginFileld.tsx
    │   │   ├── ./frontend/components/features/RecordFiled.tsx
    │   │   └── ./frontend/components/features/SignUpField.tsx
    │   └── ./frontend/components/ui
    │       ├── ./frontend/components/ui/Button.tsx
    │       ├── ./frontend/components/ui/CategorySelect.tsx
    │       ├── ./frontend/components/ui/DatePicker.tsx
    │       ├── ./frontend/components/ui/FormLine.tsx
    │       └── ./frontend/components/ui/Title.tsx
    ├── ./frontend/eslint.config.mjs
    ├── ./frontend/next-env.d.ts
    ├── ./frontend/next.config.ts
    ├── ./frontend/package-lock.json
    ├── ./frontend/package.json
    ├── ./frontend/postcss.config.mjs
    ├── ./frontend/public
    │   ├── ./frontend/public/file.svg
    │   ├── ./frontend/public/globe.svg
    │   ├── ./frontend/public/images
    │   │   └── ./frontend/public/images/char
    │   │       ├── ./frontend/public/images/char/level1.jpg
    │   │       ├── ./frontend/public/images/char/level10.jpg
    │   │       ├── ./frontend/public/images/char/level11.jpg
    │   │       ├── ./frontend/public/images/char/level12.jpg
    │   │       ├── ./frontend/public/images/char/level13.jpg
    │   │       ├── ./frontend/public/images/char/level14.jpg
    │   │       ├── ./frontend/public/images/char/level15.jpg
    │   │       ├── ./frontend/public/images/char/level16.jpg
    │   │       ├── ./frontend/public/images/char/level17.jpg
    │   │       ├── ./frontend/public/images/char/level18.jpg
    │   │       ├── ./frontend/public/images/char/level19.jpg
    │   │       ├── ./frontend/public/images/char/level2.jpg
    │   │       ├── ./frontend/public/images/char/level20.jpg
    │   │       ├── ./frontend/public/images/char/level3.jpg
    │   │       ├── ./frontend/public/images/char/level4.jpg
    │   │       ├── ./frontend/public/images/char/level5.jpg
    │   │       ├── ./frontend/public/images/char/level6.jpg
    │   │       ├── ./frontend/public/images/char/level7.jpg
    │   │       ├── ./frontend/public/images/char/level8.jpg
    │   │       └── ./frontend/public/images/char/level9.jpg
    │   ├── ./frontend/public/next.svg
    │   ├── ./frontend/public/vercel.svg
    │   └── ./frontend/public/window.svg
    ├── ./frontend/tsconfig.json
    ├── ./frontend/tsconfig.tsbuildinfo
    └── ./frontend/types
        └── ./frontend/types/api.ts

28 directories, 95 files
