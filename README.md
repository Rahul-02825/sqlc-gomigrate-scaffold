# sqlc-gomigrate-scaffold

A production-ready microservice boilerplate built with **Go**, integrating **sqlc** for compile-time type-safe database queries and **golang-migrate** for robust schema migrations against PostgreSQL.

---

## Architecture & Features

- **Type-Safe SQL**: Uses [sqlc](https://sqlc.dev/) to generate idiomatic Go code directly from raw SQL queries without ORM runtime overhead.
- **Database Versioning**: Powered by [golang-migrate](https://github.com/golang-migrate/migrate) for deterministic up/down database migrations.
- **Modular Layout**: Structured under `user-service/` for easy extension into microservice architectures.
- **Secure Configuration**: Clean separation of secrets via environment variables (`.env`).

---

## Directory Layout

```text
.
├── user-service/
│   ├── db/
│   │   ├── migrations/         # golang-migrate up/down SQL files
│   │   │   ├── 000001_create_users.up.sql
│   │   │   └── 000001_create_users.down.sql
│   │   ├── queries/            # Raw SQL queries for sqlc code generation
│   │   │   └── users.sql
│   │   └── sqlc/               # Generated type-safe Go structs and queries
│   ├── sqlc.yaml               # sqlc configuration
│   └── ...
├── .gitignore
└── README.md