# 小城非遗影像志

## 标准命令

```bash
基于 Go 实现的 HTTP Web 项目，一款业务数据管理服务，提供业务记录处理、状态查询与实时结果展示。
go test -count=1 ./...
go run ./cmd/heritage -db heritage.db -addr :8080
```

服务提供 `/api/home` 首页聚合接口，支持 `category` 查询参数；`POST /api/articles` 创建文章。
