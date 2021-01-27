## questionnaire server

![test](https://github.com/kazuki-komori/questionnaire_server/workflows/test/badge.svg)

## セットアップ

### go module セットアップ
```bash
go build
```
### マイグレーション
```bash
sql-migrate up
```

## 環境

- Golang v1.15

**データベース**
- MySQL
  
**マイグレーションツール**
- sql-migrate